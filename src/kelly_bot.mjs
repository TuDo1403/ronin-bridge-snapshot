export class BaseBot {
	constructor(startBank, bank, fee, choice, profitRate, biasThreshold, historyThreshold, switchThreshold) {
		this.startBank = startBank;
		this.bank = bank;
		this.fee = fee;
		this.choice = choice;
		this.consecutiveLosses = 0;
		this.bankHistory = [];
		this.round = 0;
		this.resultHistory = [];
		this.profitTarget = this.bank * profitRate;
		this.peak = 0;
		this.biasThreshold = biasThreshold;
		this.historyThreshold = historyThreshold;
		this.switchThreshold = switchThreshold;
		this.CHIPS = [1, 5, 10, 25, 50, 100];
		this.index = 0;
	}

	calculateCost(stake) {
		return this.fee + stake + (stake * 5) / 100;
	}

	updateResultHistory(history) {
		this.resultHistory = history;
	}

	updateHistory() {
		this.bankHistory.push(this.bank);
	}

	handleResult(win, result) {
		this.resultHistory.push(result);

		// Update consecutive losses counter
		if (win) {
			this.consecutiveLosses = 0;
		} else {
			this.consecutiveLosses += 1;
		}

		// Only consider switching if we have enough history
		if (this.resultHistory.length >= this.historyThreshold && this.consecutiveLosses >= this.switchThreshold) {
			const recentResults = this.resultHistory;
			const headsCount = recentResults.filter((r) => r).length;
			const tailsCount = recentResults.length - headsCount;
			const totalFlips = recentResults.length;
			const headsFrequency = headsCount / totalFlips;
			const tailsFrequency = tailsCount / totalFlips;

			// If there's a bias toward one outcome, choose the opposite
			if (Math.abs(headsFrequency - tailsFrequency) > this.biasThreshold) {
				// Switch to side with lower frequency (bet against the trend)
				this.choice = headsFrequency < tailsFrequency == this.choice ? false : true;
				console.log("Switching choice by bias:", this.choice ? "Heads" : "Tails");
			}
		} else if (this.consecutiveLosses > this.switchThreshold) {
			this.choice = !this.choice;
			console.log("Switching choice by consecutive loss:", this.choice ? "Heads" : "Tails");
		}
	}
}

export class KellySwitchBot extends BaseBot {
	constructor(
		startBank,
		bank,
		fee,
		choice,
		profitRate,
		biasThreshold,
		historyThreshold,
		switchThreshold,
		trailingStopRate,
		f,
		decayRate
	) {
		super(startBank, bank, fee, choice, profitRate, biasThreshold, historyThreshold, switchThreshold);

		this.trailingStopRate = trailingStopRate;
		this.f = f;
		this.decayRate = decayRate;

		console.log("Initial Bank:", this.bank);
		console.log("Fee:", this.fee);
		console.log("Choice:", this.choice ? "Heads" : "Tails");
		console.log("Profit Target:", this.profitTarget);
		console.log("Bias Threshold:", this.biasThreshold);
		console.log("History Threshold:", this.historyThreshold);
		console.log("Switch Threshold:", this.switchThreshold);
		console.log("Trailing Stop Rate:", this.trailingStopRate);
		console.log("f:", this.f);
		console.log("Decay Rate:", this.decayRate);
		console.log("Chips:", this.CHIPS);
	}

	playRound(resultBool) {
		this.peak = Math.max(this.peak, this.bank);
		if (this.peak >= this.profitTarget) {
			// Early exit if profit target reached
			return "EARLY_EXIT";
		}
		if (this.bank < this.calculateCost(this.CHIPS[0])) {
			// Bankrupt if insufficient funds
			return "BANKRUPT";
		}

		const win =
			typeof resultBool != typeof this.choice ? JSON.parse(resultBool) === this.choice : resultBool === this.choice;
		const minCost = this.calculateCost(this.CHIPS[0]);
		const prvIndex = this.index;

		if (this.bank >= minCost) {
			const decay = Math.exp(-this.decayRate * this.bank);
			const target = this.f * decay * this.bank;
			let stake = this.CHIPS[0];
			this.index = 0;

			for (let i = this.CHIPS.length - 1; i >= 0; i--) {
				if (this.CHIPS[i] <= target) {
					stake = this.CHIPS[i];
					this.index = i;
					break;
				}
			}

			const cost = this.calculateCost(stake);
			if (cost <= this.bank) {
				this.round++;
				this.bank -= cost;
				if (win) {
					this.bank += 2 * this.CHIPS[prvIndex];
				}
			}
		} else {
			// Panic
			throw new Error("Insufficient funds to place a bet");
		}

		this.handleResult(win, resultBool);
		this.updateHistory();

		return "CONTINUE";
	}
}
