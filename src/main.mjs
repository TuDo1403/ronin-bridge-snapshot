import { Bot } from "grammy";
import { ethers } from "ethers";
import dotenv from "dotenv";
import fs from "fs";
import path from "path";
// import { sendReply } from "./utils.mjs";
// import { isAuthorized } from "./auth.mjs";
// import { isAddress } from "ethers";
import { KellySwitchBot } from "./kelly_bot.mjs";
import { loadCSV, appendToCSV } from "./utils.mjs";

dotenv.config();

const outcomeHistoryPath = "./coin_flip_results_testnet.csv";

const CoinFlipperABIPath = path.resolve("./generated/abi/CoinFlipper.abi.json");
const CoinFlipperABI = JSON.parse(fs.readFileSync(CoinFlipperABIPath, "utf-8"));

const provider = new ethers.JsonRpcProvider(process.env.RONIN_TESTNET_RPC);
const coinFlipper = new ethers.Contract("0x999802ddfeba57b5e627862a0f9c23d6ea42dd20", CoinFlipperABI, provider); // testnet
const authWallet = new ethers.Wallet(process.env.AUTH_PK, provider);
console.log("Auth wallet address:", authWallet.address);

// const bot = new Bot(process.env.BOT_TOKEN);
// const authFilePath = path.resolve("src/data/auth.production.yml");

// Listen for the "Buy" event
const coinFlipInitiatedFilter = coinFlipper.filters.CoinFlipInitiated();
const coinFlipResolvedFilter = coinFlipper.filters.CoinFlipResolved();
let rollInfo = {};

const START_BANK = 1000;
const FEE = 0.5;
const TARGET_PROFIT_RATE = 1.5;

const vrfChoices = (await loadCSV(outcomeHistoryPath)).map((row) => JSON.parse(row.vrfChoice));
const headRates = vrfChoices.filter((choice) => choice == true).length / vrfChoices.length;
const tailRates = vrfChoices.filter((choice) => choice == false).length / vrfChoices.length;

let side = true;
side = headRates > tailRates ? true : false;
console.log("Side:", side ? "Heads" : "Tails");
console.log("Head Rates:", headRates);
console.log("Tail Rates:", tailRates);

// bias_threshold: 0.004395507489315027 (Range: 0.0001 to 0.5)
// history_threshold: 97.88042395892099 (Range: 1.0 to 100.0)
// switch_threshold: 1.100139769823249 (Range: 1.0 to 500.0)
// profit_rate: 4.993780144643149 (Range: 1.5 to 5.0)
// trailing_stop_rate: 0.8773144251216427 (Range: 0.5 to 1.0)
// f: 0.39638644632467956 (Range: 0.001 to 0.5)
// decay_rate: 0.000586849587412405 (Range: 1e-06 to 0.1)
const kellyBot = new KellySwitchBot(
	START_BANK,
	START_BANK,
	FEE,
	side,
	TARGET_PROFIT_RATE,
	0.004395507489315027,
	97.88042395892099,
	1,
	0.6642856515983583,
	0.39638644632467956,
	0.000586849587412405
);
kellyBot.updateResultHistory(vrfChoices);

// event CoinFlipInitiated(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool choice_, uint256 nftId_)
coinFlipper.on(coinFlipInitiatedFilter, async (event) => {
	const { args } = event;
	const player = args[0];
	const configId = args[1];
	const reqHash = args[2];
	const choice = args[3];
	const nftId = args[4];

	console.log(
		`${player} initiated a coin flip with configId: ${configId}, reqHash: ${reqHash}, choice: ${choice}, nftId: ${nftId}`
	);

	rollInfo[reqHash] = {
		player: player,
		configId: configId,
		choice: choice,
		nftId: nftId,
		won: false,
	};

	// sendReply(ctx, message);
	// bot.api.sendMessage(chatId, message).catch((err) => {
	// 	console.error("Error sending message:", err);
	// });
	// console.log("Message sent!");
});

// event CoinFlipResolved(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool playerWin_)
coinFlipper.on(coinFlipResolvedFilter, async (event) => {
	const { args } = event;
	const player = args[0];
	const configId = args[1];
	const reqHash = args[2];
	const playerWin = args[3];

	console.log(
		`${player} resolved a coin flip with configId: ${configId}, reqHash: ${reqHash}, playerWin: ${playerWin}`
	);

	let vrfChoice = false; // true = head, false = tail
	if (rollInfo[reqHash].choice) {
		// true = head, false = tail
		if (playerWin) {
			vrfChoice = true;
		}
	} else {
		// false = tail
		if (!playerWin) {
			vrfChoice = true;
		}
	}

	await appendToCSV(outcomeHistoryPath, [
		{
			choice: rollInfo[reqHash].choice ? "true" : "false",
			bet_amount: 0,
			won: playerWin ? "true" : "false",
			vrfChoice: vrfChoice ? "true" : "false",
		},
	]);

	if (player == authWallet.address) {
		const action = kellyBot.playRound(vrfChoice);
		const choice = kellyBot.choice;
		console.log("KellyBot's bank:", kellyBot.bank, "Round:", kellyBot.round);

		if (action == "EARLY_EXIT" || action == "BANKRUPT") {
			if (action == "EARLY_EXIT") {
				console.log("KellyBot reached profit target, exiting...");
				return;
			}
			if (action == "BANKRUPT") {
				console.log("KellyBot is bankrupt, exiting...");
				return;
			}
		}

		const amount = ethers.parseEther(kellyBot.CHIPS[kellyBot.index].toString());
		const { vrfFee, fee, gasPrice } = await estimateFee(amount);
		console.log("Estimated fee:", fee + amount);
		let gas = await coinFlipper.flipACoin.estimateGas(choice, kellyBot.index, 0, {
			from: authWallet.address,
			value: fee + amount,
		});

		const receipt = await coinFlipper.connect(authWallet).flipACoin(choice, kellyBot.index, 0, {
			value: fee + amount,
		});

		await receipt.wait();
		console.log("Transaction hash:", receipt.hash);

		const newFee = ethers.formatEther((gas * gasPrice) / BigInt(1e18) + vrfFee);
		if (kellyBot.fee != newFee) {
			console.log("Fee changed, updating...", kellyBot.fee, "->", newFee);
			kellyBot.fee = parseFloat(newFee);
		}
	}

	// bot.api.sendMessage(chatId, message).catch((err) => {
	// 	console.error("Error sending message:", err);
	// });
});

const estimateFee = async (amount) => {
	const gasPrice = (await provider.getFeeData()).gasPrice;
	let { estFee } = await coinFlipper.estimateFee(gasPrice);
	// buffer 10% for gas
	const vrfFee = BigInt(estFee) + (estFee * BigInt(10)) / BigInt(100);
	const fee = vrfFee + (amount * BigInt(5)) / BigInt(100);
	return { vrfFee, fee, gasPrice };
};

console.log("Initial Triggering...");

const amount = ethers.parseEther(kellyBot.CHIPS[5].toString());
const { vrfFee, fee, gasPrice } = await estimateFee(amount);
const gas = await coinFlipper.flipACoin.estimateGas(kellyBot.choice, kellyBot.index, 0, {
	from: authWallet.address,
	value: fee + amount,
});

const receipt = await coinFlipper.connect(authWallet).flipACoin(kellyBot.choice, kellyBot.index, 0, {
	value: fee + amount,
});

await receipt.wait();
console.log("Transaction hash:", receipt.hash);

const newFee = ethers.formatEther((gas * gasPrice) / BigInt(1e18) + vrfFee);
if (kellyBot.fee != newFee) {
	console.log("Fee changed, updating...", kellyBot.fee, "->", newFee);
	kellyBot.fee = parseFloat(newFee);
}
