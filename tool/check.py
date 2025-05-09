import math
import matplotlib.pyplot as plt
import pandas as pd
# import csv, math, ace_tools as tools

file_path = "coin_flip_results.csv"
chips = [1, 5, 10, 25, 50, 100]
start_bank = 1000.0
kellyF = 0.17


class KellyLadderBot:
    """
    Hybrid: stake palette = chips ladder,
    starting rung chosen by current Kelly target,
    climbs ladder on each loss up to max rung,
    resets rung after a win.
    Switches coin side after 2 consecutive losses.
    """

    def __init__(self, bank, fee, f):
        self.bank = bank
        self.fee = fee
        self.f = f
        self.index = 0
        self.round = 0
        self.choice = True  # True=head
        self.loss_streak = 0
        self.bank_history = []

    def next_rung_from_kelly(self):
        target = self.f * self.bank
        rung = 0
        for i, c in enumerate(chips):
            if c <= target:
                rung = i
            else:
                break
        return rung

    def play_round(self, result_bool):
        # Update index against new Kelly target if lower than current
        desired = self.next_rung_from_kelly()
        if self.index < desired:
            self.index = desired  # grow with bank automatically
        stake = chips[self.index]
        cost = stake + self.fee
        if cost > self.bank:
            self.bank_history.append(self.bank)
            return
        win = result_bool == self.choice
        self.round += 1
        self.bank -= cost
        if win:
            self.bank += stake * 2
            # reset rung to Kelly optimum after win
            self.index = self.next_rung_from_kelly()
            self.loss_streak = 0
        else:
            self.loss_streak += 1
            if self.index < len(chips) - 1:
                self.index += 1  # climb
        # side switching rule
        if self.loss_streak >= 2 and not win:
            self.choice = not self.choice
            self.loss_streak = 0
        self.bank_history.append(self.bank)


# baseline LadderSwitch for comparison
class LadderSwitchBot:
    def __init__(self, bank, fee):
        self.bank = bank
        self.fee = fee
        self.index = 0
        self.round = 0
        self.choice = True
        self.loss_streak = 0
        self.bank_history = []

    def play_round(self, result_bool):
        stake = chips[self.index]
        cost = stake + self.fee
        if cost > self.bank:
            self.bank_history.append(self.bank)
            return
        win = result_bool == self.choice
        self.round += 1
        self.bank -= cost
        if win:
            self.bank += stake * 2
            self.index = 0
            self.loss_streak = 0
        else:
            self.loss_streak += 1
            if self.index < len(chips) - 1:
                self.index += 1
        if self.loss_streak >= 2 and not win:
            self.choice = not self.choice
            self.loss_streak = 0
        self.bank_history.append(self.bank)


kelly_ladder = KellyLadderBot(start_bank, 0.0, kellyF)
ladder_switch = LadderSwitchBot(start_bank, 0.0)
data = pd.read_csv(file_path)
for _, row in data.iterrows():
    result_bool = row['vrfChoice'] == "true"
    kelly_ladder.play_round(result_bool)
    ladder_switch.play_round(result_bool)


def insolvency_round(banks):
    for i, b in enumerate(banks):
        if b < chips[0]:
            return i + 1
    return None


flips = len(kelly_ladder.bank_history)
summary = pd.DataFrame(
    {
        "Bot": ["KellyLadderHybrid", "LadderSwitch"],
        "Final Bank": [kelly_ladder.bank, ladder_switch.bank],
        "Profit": [kelly_ladder.bank - start_bank, ladder_switch.bank - start_bank],
        "Flips Played": [flips, flips],
        "Reached Insolvency": [
            insolvency_round(kelly_ladder.bank_history) is not None,
            insolvency_round(ladder_switch.bank_history) is not None,
        ],
        "CAGR %": [
            (
                (math.pow(kelly_ladder.bank / start_bank, 1 / flips) - 1) * 100
                if kelly_ladder.bank > 0
                else -100.0
            ),
            (
                (math.pow(ladder_switch.bank / start_bank, 1 / flips) - 1) * 100
                if ladder_switch.bank > 0
                else -100.0
            ),
        ],
    }
)

# tools.display_dataframe_to_user("Hybrid vs LadderSwitch Summary", summary)

# Plot
plt.figure()
plt.plot(kelly_ladder.bank_history, label="KellyLadderHybrid")
plt.plot(ladder_switch.bank_history, label="LadderSwitch")
plt.xlabel("Flip #")
plt.ylabel("Bankroll")
plt.title("Kelly‑Ladder Hybrid vs LadderSwitch")
plt.legend()
plt.tight_layout()
plt.savefig("coin_flip_results.png")
