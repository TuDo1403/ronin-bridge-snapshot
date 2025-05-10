import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
import math

# Configuration
FILE_PATH = "coin_flip_results.csv"
CHIPS = [1, 5, 10, 25, 50, 100]
START_BANK = 1000
KELLY_FRACTION = 0.17
FEE = 0.0
SWITCH_THRESHOLD = 3  # Switch after this many consecutive losses


def run_simulation(shuffle=False):
    kelly_bot = KellySwitchBot(START_BANK, FEE, KELLY_FRACTION, True)
    ladder_bot = LadderSwitchBot(START_BANK, FEE, True)

    # Read data using pandas
    data = pd.read_csv(FILE_PATH)
    key = "vrfChoice"
    results = data[key].values
    if shuffle:
        np.random.shuffle(results)
    # Convert results to DataFrame
    data = pd.DataFrame(results, columns=[key])

    # Iterate through the DataFrame rows
    for _, row in data.iterrows():
        result_bool = row[key]
        kelly_bot.play_round(result_bool)
        ladder_bot.play_round(result_bool)

    return kelly_bot, ladder_bot


def run_multiple_simulations(n_runs=100):
    # Track survival statistics
    kelly_survivals = []
    ladder_survivals = []
    kelly_wins = 0

    # Read data using pandas
    data = pd.read_csv(FILE_PATH)
    key = "vrfChoice"
    results = data[key].values

    for i in range(n_runs):
        # Create new bots for each run
        kelly_bot = KellySwitchBot(START_BANK, FEE, KELLY_FRACTION, True)
        ladder_bot = LadderSwitchBot(START_BANK, FEE, True)

        # Shuffle results if requested
        run_results = np.random.permutation(results) if i > 0 else results

        # Run simulation until one or both bots go insolvent
        for round_num, result in enumerate(run_results, 1):
            if kelly_bot.bank >= CHIPS[0]:
                kelly_bot.play_round(result)

            if ladder_bot.bank >= CHIPS[0]:
                ladder_bot.play_round(result)

            # If both are insolvent, end this simulation run
            if kelly_bot.bank < CHIPS[0] and ladder_bot.bank < CHIPS[0]:
                break

        # Record how many rounds each bot survived
        kelly_rounds = kelly_bot.round
        ladder_rounds = ladder_bot.round

        kelly_survivals.append(kelly_rounds)
        ladder_survivals.append(ladder_rounds)

        if kelly_bot.bank > ladder_bot.bank:
            kelly_wins += 1

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
            "Win Rate": [kelly_wins / n_runs, (n_runs - kelly_wins) / n_runs],
            "Profit Rate": [
                (np.mean(kelly_survivals) - START_BANK) / START_BANK,
                (np.mean(ladder_survivals) - START_BANK) / START_BANK,
            ],
        }
    )

    return summary, kelly_survivals, ladder_survivals


class BaseBot:
    def __init__(self, bank, fee, choice):
        self.bank = bank
        self.fee = fee
        self.choice = choice
        self.consecutive_losses = 0
        self.bank_history = []
        self.round = 0
        self.result_history = []

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
            len(self.result_history) >= 20
            and self.consecutive_losses >= SWITCH_THRESHOLD
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
            if abs(heads_frequency - tails_frequency) > 0.15:  # 15% threshold
                # Switch to the side with lower frequency (bet against the trend)
                self.choice = (
                    heads_frequency > tails_frequency
                )  # True if should choose tails

                print(
                    f"Switching choice at round {self.round}: heads_frequency={heads_frequency:.2f}, tails_frequency={tails_frequency:.2f} -> choice={self.choice}"
                )

            # else:
            #     self.choice = not self.choice  # Switch back to the original choice


class KellySwitchBot(BaseBot):
    def __init__(self, bank, fee, f, choice):
        super().__init__(bank, fee, choice)
        self.f = f

    def play_round(self, result_bool):
        win = result_bool == self.choice
        min_cost = CHIPS[0] + self.fee

        if self.bank >= min_cost:
            # target = self.f * self.bank
            decay = math.exp(-0.002 * self.bank)  # λ = 0.002
            target = self.f * decay * self.bank
            stake = CHIPS[0]
            for c in reversed(CHIPS):
                if c <= target:
                    stake = c
                    break

            cost = stake + self.fee
            if cost <= self.bank:
                self.round += 1
                self.bank -= cost
                if win:
                    self.bank += stake * 2

        self.handle_result(win, result_bool)
        self.update_history()


class LadderSwitchBot(BaseBot):
    def __init__(self, bank, fee, choice):
        super().__init__(bank, fee, choice)
        self.index = 0

    def play_round(self, result_bool):
        win = result_bool == self.choice
        stake = CHIPS[self.index]
        cost = stake + self.fee

        if cost <= self.bank:
            self.round += 1
            self.bank -= cost
            if win:
                self.bank += stake * 2
                self.index = 0
            else:
                if self.consecutive_losses > SWITCH_THRESHOLD:
                    # If we have lost 3 times in a row, switch to the opposite side
                    self.choice = not self.choice

                # Apply late stage decay - if bank is getting low, be more conservative
                if (
                    self.bank < START_BANK * 0.2
                ):  # When bank falls below 20% of starting value
                    # Either stay at current level or decrease it
                    if self.index > 0 and self.consecutive_losses > 3:
                        self.index -= 1  # Reduce bet size when losing and low on funds

                    # Normal ladder climbing behavior
                    elif self.index < len(CHIPS) - 1:
                        self.index += 1  # climb ladder

        self.handle_result(win, result_bool)
        self.update_history()


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

    # Run multiple simulations and generate summary
    summary, kelly_survivals, ladder_survivals = run_multiple_simulations(n_runs=100)
    print("\n=== Multiple Simulations Summary ===")
    print(summary.to_string(index=False))
    print("\n=== Survival Statistics ===")
    print(
        f"KellySwitch: {np.mean(kelly_survivals):.2f} ± {np.std(kelly_survivals):.2f}"
    )
    print(
        f"LadderSwitch: {np.mean(ladder_survivals):.2f} ± {np.std(ladder_survivals):.2f}"
    )
    plt.figure(figsize=(10, 6))
    plt.hist(kelly_survivals, bins=15, alpha=0.5, label="KellySwitch")
    plt.hist(ladder_survivals, bins=15, alpha=0.5, label="LadderSwitch")
    plt.xlabel("Rounds Survived")
    plt.ylabel("Frequency")
    plt.title("Survival Histogram")
    plt.legend()
    plt.grid(True, alpha=0.3)
    plt.tight_layout()
    plt.show()
