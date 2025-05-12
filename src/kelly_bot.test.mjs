import { KellySwitchBot } from "./kelly_bot.mjs";
import { loadCSV } from "./utils.mjs";

const FILE_PATH = "coin_flip_results_testnet.csv";
const START_BANK = 1000;
const FEE = 0.6;
const SIDE = true; // true = head, false = tail
const TARGET_PROFIT_RATE = 1.5;

const calculateCagr = (initialValue, finalValue, periods) => {
	if (periods <= 0) return -100;
	return (Math.pow(finalValue / initialValue, 1 / periods) - 1) * 100;
};

const generateSummary = (bot) => {
	const flips = bot.bankHistory.length;
	console.log("Final Bank:", bot.bank.toFixed(2));
	console.log("Flips Played:", flips);
	console.log("CAGR %:", calculateCagr(START_BANK, bot.bank, flips).toFixed(2));
};

// bias_threshold: 0.030708025079715994 (Range: 0.0001 to 0.5)
// history_threshold: 25.129938563763393 (Range: 1.0 to 100.0)
// switch_threshold: 492.70801784798215 (Range: 1.0 to 500.0)
// profit_rate: 1.5587565593751913 (Range: 1.5 to 5.0)
// trailing_stop_rate: 0.6642856515983583 (Range: 0.5 to 1.0)
// f: 0.4999081120383311 (Range: 0.001 to 0.5)
// decay_rate: 8.818522269623361e-06 (Range: 1e-06 to 0.1)
let kellyBot = new KellySwitchBot(
	START_BANK,
	START_BANK,
	FEE,
	SIDE,
	TARGET_PROFIT_RATE,
	0.030708025079715994,
	25.129938563763393,
	492.70801784798215,
	0.6642856515983583,
	0.4999081120383311,
	8.818522269623361e-6
);

// startBank = 1000,
// bank = 1000,
// fee = 0.32,
// choice = true,
// profitRate = 1.7,
// biasThreshold = 0.15,
// historyThreshold = 20,
// switchThreshold = 3,
// trailingStopRate = 0.9,
// f = 0.17,
// decayRate = 0.002

const vrfChoices = (await loadCSV(FILE_PATH)).map((row) => row.vrfChoice);
vrfChoices.forEach((choice) => {
	kellyBot.playRound(choice);
	console.log("Kelly:", "Choice:", kellyBot.choice ? "True" : "False", "Bank:", kellyBot.bank);
});

generateSummary(kellyBot);
