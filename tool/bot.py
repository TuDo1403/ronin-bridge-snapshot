import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
import math

# Configuration
FILE_PATH = "coin_flip_results_testnet.csv"
CHIPS = [1, 5, 10, 25, 50, 100]
START_BANK = 1000
FEE = 0.6
SIDE = True
KEY = "vrfChoice"  # Column name in the CSV file
FUZZ = False
target_profit_rate = 1.5
print("Expected profit:", target_profit_rate * START_BANK - START_BANK)


# Best Parameters:
# bias_threshold 0: 0.05917482528730716 (Range: 0.0001 to 0.5)
# history_threshold 1: 1.7686155374229706 (Range: 1.0 to 1000.0)
# switch_threshold 2: 4.8800990267340705 (Range: 1.0 to 500.0)
# stage_decay 3: 0.07740252756707897 (Range: 0.05 to 0.9)
# profit_rate 4: 1.2162022723649835 (Range: 1.2 to 5.0)
# trailing_stop_rate 5: 0.21764660190214433 (Range: 0.0 to 1.0)
# Best Objective Values:
# Objective 0 (Rounds Played): 447.64
# Objective 1 (Profit): -966.1051999999987
# Best Solution:
# Rounds Played: 447.64
# Profit: -966.1051999999987


def run_simulation(shuffle=False):
    # bias_threshold: 0.004395507489315027 (Range: 0.0001 to 0.5)
    # history_threshold: 97.88042395892099 (Range: 1.0 to 100.0)
    # switch_threshold: 1.100139769823249 (Range: 1.0 to 500.0)
    # profit_rate: 4.993780144643149 (Range: 1.5 to 5.0)
    # trailing_stop_rate: 0.8773144251216427 (Range: 0.5 to 1.0)
    # f: 0.39638644632467956 (Range: 0.001 to 0.5)
    # decay_rate: 0.000586849587412405 (Range: 1e-06 to 0.1)
    kelly_bot = KellySwitchBot(
        start_bank=START_BANK,
        bank=START_BANK,
        fee=FEE,
        choice=SIDE,
        profit_rate=target_profit_rate,
        bias_threshold=0.004395507489315027,
        history_threshold=97.88042395892099,
        switch_threshold=1.100139769823249,
        f=0.39638644632467956,
        decay_rate=0.000586849587412405,
    )

    # bias_threshold: 0.058354658220187185 (Range: 0.0001 to 0.5)
    # history_threshold: 4.410802301076179 (Range: 1.0 to 100.0)
    # switch_threshold: 5.26498079554969 (Range: 1.0 to 500.0)
    # stage_decay: 0.0789483282295336 (Range: 0.05 to 0.9)
    # profit_rate: 1.5186841511040372 (Range: 1.5 to 5.0)
    # trailing_stop_rate: 0.056903523359389335 (Range: 0.0 to 1.0)
    ladder_bot = LadderSwitchBot(
        start_bank=START_BANK,
        bank=START_BANK,
        fee=FEE,
        choice=SIDE,
        profit_rate=target_profit_rate,
        bias_threshold=0.058354658220187185,
        history_threshold=4.410802301076179,
        switch_threshold=5.26498079554969,
        stage_decay=0.0789483282295336,
        trailing_stop_rate=0.056903523359389335,
    )

    # Read data using pandas
    data = pd.read_csv(FILE_PATH)
    key = KEY
    results = data[key].values
    if shuffle:
        np.random.shuffle(results)
    # Convert results to DataFrame
    data = pd.DataFrame(results, columns=[key])

    # Iterate through the DataFrame rows
    for _, row in data.iterrows():
        result_bool = row[key]
        kelly_bot.play_round(result_bool)
        print("Kelly:", "Choice:", kelly_bot.choice, "Bank:", kelly_bot.bank)
        ladder_bot.play_round(result_bool)

    return kelly_bot, ladder_bot


def run_multiple_simulations(n_runs=100):
    # Track survival statistics
    kelly_survivals = []
    ladder_survivals = []
    kelly_wins = 0
    ladder_wins = 0
    kelly_banks = []
    ladder_banks = []
    kelly_exits = []  # Initialize with dummy value
    ladder_exits = []  # Initialize with dummy value

    for i in range(n_runs):
        # bias_threshold: 0.004395507489315027 (Range: 0.0001 to 0.5)
        # history_threshold: 97.88042395892099 (Range: 1.0 to 100.0)
        # switch_threshold: 1.100139769823249 (Range: 1.0 to 500.0)
        # profit_rate: 4.993780144643149 (Range: 1.5 to 5.0)
        # trailing_stop_rate: 0.8773144251216427 (Range: 0.5 to 1.0)
        # f: 0.39638644632467956 (Range: 0.001 to 0.5)
        # decay_rate: 0.000586849587412405 (Range: 1e-06 to 0.1)
        kelly_bot = KellySwitchBot(
            start_bank=START_BANK,
            bank=START_BANK,
            fee=FEE,
            choice=SIDE,
            profit_rate=target_profit_rate,
            bias_threshold=0.004395507489315027,
            history_threshold=25.88042395892099,
            switch_threshold=3.100139769823249,
            f=0.39638644632467956,
            decay_rate=0.000586849587412405,
        )

        # bias_threshold: 0.058354658220187185 (Range: 0.0001 to 0.5)
        # history_threshold: 4.410802301076179 (Range: 1.0 to 100.0)
        # switch_threshold: 5.26498079554969 (Range: 1.0 to 500.0)
        # stage_decay: 0.0789483282295336 (Range: 0.05 to 0.9)
        # profit_rate: 1.5186841511040372 (Range: 1.5 to 5.0)
        # trailing_stop_rate: 0.056903523359389335 (Range: 0.0 to 1.0)
        ladder_bot = LadderSwitchBot(
            start_bank=START_BANK,
            bank=START_BANK,
            fee=FEE,
            choice=SIDE,
            profit_rate=target_profit_rate,
            bias_threshold=0.058354658220187185,
            history_threshold=4.410802301076179,
            switch_threshold=5.26498079554969,
            stage_decay=0.0789483282295336,
            trailing_stop_rate=0.056903523359389335,
        )

        # Shuffle results if requested
        # np.random.seed(i)  # Set seed for reproducibility
        run_results = np.random.choice([True, False], size=1000)
        kelly_early_exit = False
        ladder_early_exit = False

        # Run simulation until one or both bots go insolvent
        for round_num, result in enumerate(run_results, 1):
            if kelly_bot.bank >= CHIPS[0] and (
                len(kelly_exits) == 0
                or kelly_exits[len(kelly_exits) - 1]["round"] != round_num
            ):
                out = kelly_bot.play_round(result)
                if out == "EARLY_EXIT" or out == "BANKRUPT":
                    kelly_exits.append({"run": i, "round": round_num})
                    kelly_early_exit = True

            if ladder_bot.bank >= CHIPS[0] and (
                len(ladder_exits) == 0
                or ladder_exits[len(ladder_exits) - 1]["round"] != round_num
            ):
                out = ladder_bot.play_round(result)
                if out == "EARLY_EXIT" or out == "BANKRUPT":
                    ladder_exits.append({"run": i, "round": round_num})
                    ladder_early_exit = True

            # If both are insolvent, end this simulation run
            if (
                kelly_bot.bank < CHIPS[0] + kelly_bot.fee
                and ladder_bot.bank < CHIPS[0] + ladder_bot.fee
            ) or (kelly_early_exit and ladder_early_exit):
                break

        # Record how many rounds each bot survived
        kelly_rounds = kelly_bot.round
        ladder_rounds = ladder_bot.round
        kelly_banks.append(kelly_bot.bank)
        ladder_banks.append(ladder_bot.bank)

        kelly_survivals.append(kelly_rounds)
        ladder_survivals.append(ladder_rounds)

        if kelly_bot.bank > START_BANK:
            kelly_wins += 1
        elif ladder_bot.bank > START_BANK:
            ladder_wins += 1

        print(
            f"Run {i + 1}/{n_runs}: KellySwitch Bank {kelly_bot.bank}, LadderSwitch Bank {ladder_bot.bank}"
        )

    # Create summary DataFrame
    summary = pd.DataFrame(
        {
            "Bot": ["KellySwitch", "LadderSwitch"],
            "Avg Rounds Survived": [
                np.mean(kelly_survivals),
                np.mean(ladder_survivals),
            ],
            "Max Rounds Survived": [np.max(kelly_survivals), np.max(ladder_survivals)],
            "Min Rounds Survived": [np.min(kelly_survivals), np.min(ladder_survivals)],
            "Win Rate": [kelly_wins / n_runs, ladder_wins / n_runs],
            "Profit Rate": [
                (np.mean(kelly_banks) / START_BANK),
                (np.mean(ladder_banks) / START_BANK),
            ],
            "Early Exit Rate": [
                len(kelly_exits) / n_runs,
                len(ladder_exits) / n_runs,
            ],
        }
    )

    return summary, kelly_banks, ladder_banks


class BaseBot:
    def __init__(
        self,
        start_bank,
        bank,
        fee,
        choice,
        profit_rate,
        bias_threshold,
        history_threshold,
        switch_threshold,
    ):
        self.start_bank = start_bank
        self.bank = bank
        self.fee = fee
        self.choice = choice
        self.consecutive_losses = 0
        self.bank_history = []
        self.round = 0
        self.result_history = []
        self.profit_target = self.bank * profit_rate
        self.peak = 0
        self.bias_threshold = bias_threshold
        self.history_threshold = history_threshold
        self.switch_threshold = switch_threshold
        self.CHIPS = [1, 5, 10, 25, 50, 100]

    def calculate_fee(self, stake):
        return self.fee + stake + stake * 5 / 100

    def update_history(self):
        self.bank_history.append(self.bank)

    def handle_result(self, win, result):
        self.result_history.append(result)

        # Update consecutive losses counter
        if win:
            self.consecutive_losses = 0
        else:
            self.consecutive_losses += 1

        # Only consider switching if we have enough history
        if (
            len(self.result_history) >= self.history_threshold
            and self.consecutive_losses >= self.switch_threshold
        ):
            # Calculate recent results
            recent_results = self.result_history
            heads_count = sum(1 for r in recent_results if r)
            tails_count = len(recent_results) - heads_count

            # Calculate outcome frequencies
            total_flips = len(recent_results)
            heads_frequency = heads_count / total_flips
            tails_frequency = tails_count / total_flips

            # If there's a bias toward one outcome, choose the opposite
            # (i.e., if we see more heads, bet on tails and vice versa)
            if (
                abs(heads_frequency - tails_frequency) > self.bias_threshold
            ):  # 15% threshold
                # Switch to the side with lower frequency (bet against the trend)
                # self.choice = (
                #     heads_frequency < tails_frequency
                # )  # True if should choose tails
                self.choice = (
                    False if heads_frequency > tails_frequency == self.choice else True
                )

        elif self.consecutive_losses > self.switch_threshold:
            self.choice = not self.choice  # Switch back to the original choice


class KellySwitchBot(BaseBot):
    def __init__(
        self,
        start_bank=1000,
        bank=1000,
        fee=0.32,
        choice=True,
        profit_rate=1.7,
        bias_threshold=0.15,
        history_threshold=20,
        switch_threshold=3,
        trailing_stop_rate=0.9,
        f=0.17,
        decay_rate=0.002,
    ):
        super().__init__(
            start_bank,
            bank,
            fee,
            choice,
            profit_rate,
            bias_threshold,
            history_threshold,
            switch_threshold,
        )
        self.f = f
        self.trailing_stop_rate = trailing_stop_rate
        self.decay_rate = decay_rate

    def play_round(self, result_bool):
        self.peak = max(self.peak, self.bank)
        if self.peak >= self.profit_target:
            # If we have reached the profit target, stop playing
            return "EARLY_EXIT"
        if self.bank < self.calculate_fee(self.CHIPS[0]):
            # If we don't have enough to play, stop playing
            return "BANKRUPT"

        win = result_bool == self.choice
        min_cost = self.calculate_fee(self.CHIPS[0])

        if self.bank >= min_cost:
            # target = self.f * self.bank
            decay = math.exp(-self.decay_rate * self.bank)  # λ = 0.002
            target = self.f * decay * self.bank
            stake = self.CHIPS[0]
            for c in reversed(self.CHIPS):
                if c <= target:
                    stake = c
                    break

            cost = self.calculate_fee(stake)
            if cost <= self.bank:
                self.round += 1
                self.bank -= cost
                if win:
                    self.bank += stake * 2

        self.handle_result(win, result_bool)
        self.update_history()

        return "CONTINUE"


class LadderSwitchBot(BaseBot):
    def __init__(
        self,
        start_bank=1000,
        bank=1000,
        fee=0.32,
        choice=True,
        profit_rate=1.7,
        bias_threshold=0.15,
        history_threshold=20,
        switch_threshold=3,
        stage_decay=0.2,
        trailing_stop_rate=0.9,
    ):
        super().__init__(
            start_bank,
            bank,
            fee,
            choice,
            profit_rate,
            bias_threshold,
            history_threshold,
            switch_threshold,
        )
        self.index = 0
        self.stage_decay = stage_decay
        self.trailing_stop_rate = trailing_stop_rate

    def play_round(self, result_bool):
        self.peak = max(self.peak, self.bank)
        if self.peak >= self.profit_target:
            # If we have reached the profit target, stop playing
            return "EARLY_EXIT"
        if self.bank < self.calculate_fee(self.CHIPS[0]):
            # If we don't have enough to play, stop playing
            return "BANKRUPT"

        win = result_bool == self.choice
        stake = self.CHIPS[self.index]
        cost = self.calculate_fee(stake)
        prv_idx = self.index

        if cost <= self.bank:
            self.round += 1
            self.bank -= cost
            if win:
                self.bank += self.CHIPS[prv_idx] * 2
                self.index = 0
            else:
                # Apply late stage decay - if bank is getting low, be more conservative
                if (
                    self.bank < self.start_bank * self.stage_decay
                ):  # When bank falls below 20% of starting value
                    # Either stay at current level or decrease it
                    if (
                        self.index > 0
                        and self.consecutive_losses > self.switch_threshold
                    ):
                        self.index -= 1  # Reduce bet size when losing and low on funds

                # Normal ladder climbing behavior
                elif self.index < len(self.CHIPS) - 1:
                    self.index += 1  # climb ladder

        self.handle_result(win, result_bool)
        self.update_history()

        return "CONTINUE"


def insolvency_round(banks):
    for i, b in enumerate(banks):
        if b < CHIPS[0]:
            return i + 1
    return None


def calculate_cagr(final_value, initial_value, periods):
    if final_value <= 0:
        return -100.0
    return (math.pow(final_value / initial_value, 1 / periods) - 1) * 100


def generate_summary(kelly_bot, ladder_bot):
    flips = len(kelly_bot.bank_history)

    summary = pd.DataFrame(
        {
            "Bot": ["KellySwitch", "LadderSwitch"],
            "Final Bank": [kelly_bot.bank, ladder_bot.bank],
            "Profit": [kelly_bot.bank - START_BANK, ladder_bot.bank - START_BANK],
            "Flips Played": [flips, flips],
            "Reached Insolvency At Flip": [
                insolvency_round(kelly_bot.bank_history),
                insolvency_round(ladder_bot.bank_history),
            ],
            "CAGR %": [
                calculate_cagr(kelly_bot.bank, START_BANK, flips),
                calculate_cagr(ladder_bot.bank, START_BANK, flips),
            ],
        }
    )

    return summary


def plot_results(kelly_bot, ladder_bot):
    plt.figure(figsize=(10, 6))
    plt.plot(kelly_bot.bank_history, label="KellySwitch")
    plt.plot(ladder_bot.bank_history, label="LadderSwitch")

    # Add horizontal line for starting bank balance
    plt.axhline(
        y=START_BANK, color="r", linestyle="--", label=f"Starting Bank ({START_BANK})"
    )

    plt.xlabel("Flip #")
    plt.ylabel("Bankroll")
    plt.title("Bankroll Trajectory – Dynamic Side & Hybrid Sizing")
    plt.legend()
    plt.grid(True, alpha=0.3)
    plt.tight_layout()
    plt.show()


if __name__ == "__main__":
    kelly_bot, ladder_bot = run_simulation(shuffle=False)
    summary = generate_summary(kelly_bot, ladder_bot)

    print("\n=== Hybrid vs. Ladder‑Switch Summary ===")
    print(summary.to_string(index=False))

    plot_results(kelly_bot, ladder_bot)

    if FUZZ:
        # Run multiple simulations and generate summary
        summary, kelly_banks, ladder_banks = run_multiple_simulations(n_runs=1_000)
        print("\n=== Multiple Simulations Summary ===")
        print(summary.to_string(index=False))
        print("\n=== Profit Statistics ===")
        print(f"KellySwitch: {np.mean(kelly_banks):.2f} ± {np.std(kelly_banks):.2f}")
        print(f"LadderSwitch: {np.mean(ladder_banks):.2f} ± {np.std(ladder_banks):.2f}")
        plt.figure(figsize=(10, 6))
        plt.hist(kelly_banks, bins=15, alpha=0.5, label="KellySwitch")
        plt.hist(ladder_banks, bins=15, alpha=0.5, label="LadderSwitch")
        plt.xlabel("Rounds")
        plt.ylabel("Profit")
        plt.title("Survival Histogram")
        plt.legend()
        plt.grid(True, alpha=0.3)
        plt.tight_layout()
        plt.show()
