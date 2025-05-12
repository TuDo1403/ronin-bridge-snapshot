import fs from "fs/promises";
import { parse, stringify } from "csv/sync";

/**
 * Loads existing CSV data and appends new rows
 * @param {string} filePath - Path to the CSV file
 * @param {Array<{choice: number, bet_amount: number, won: boolean, vrfChoice: boolean}>} newRows - Array of new rows to append
 */
export async function appendToCSV(filePath, newRows) {
	try {
		let existingData = [];

		// Try to read existing file
		try {
			const content = await fs.readFile(filePath, "utf-8");
			existingData = parse(content, {
				columns: true,
				skip_empty_lines: true,
			});
		} catch (err) {
			// If file doesn't exist, we'll create it with headers
			existingData = [];
		}

		// Combine existing data with new rows
		const allData = [...existingData, ...newRows];

		// Convert to CSV string
		const csvString = stringify(allData, {
			header: true,
			columns: ["choice", "bet_amount", "won", "vrfChoice"],
		});

		// Write back to file
		await fs.writeFile(filePath, csvString);
	} catch (error) {
		console.error("Error handling CSV:", error);
		throw error;
	}
}

/**
 * Loads and returns data from a CSV file
 * @param {string} filePath - Path to the CSV file
 * @returns {Promise<Array<{choice: number, bet_amount: number, won: boolean, vrfChoice: boolean}>>}
 */
export async function loadCSV(filePath) {
	try {
		const content = await fs.readFile(filePath, "utf-8");
		return parse(content, {
			columns: true,
			skip_empty_lines: true,
			cast: true,
		});
	} catch (error) {
		console.error("Error loading CSV:", error);
		throw error;
	}
}
/**
 * Gets all VRF choices from the CSV data
 * @param {Array<{vrfChoice: boolean}>} data - Array of data rows
 * @returns {Array<boolean>} Array of VRF choices
 */
export function getVRFChoices(data) {
    return data.map(row => row.vrfChoice);
}