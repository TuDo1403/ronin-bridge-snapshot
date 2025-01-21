import pandas as pd  # type: ignore


# Get Sum quantity in all rows in A
# Helper function to convert string to float
def str_to_float(s: str) -> float:
    try:
        return float(s)
    except ValueError:
        return 0.0


# File paths
file_b = "ronin-weth-erc20-tx-hashes.csv"
file_a = "ronin-weth-request-withdrawals.csv"
file_c = "mainchain-weth-withdrawals.csv"

# Load CSV files into DataFrames
df_a = pd.read_csv(file_a)
df_b = pd.read_csv(file_b)
df_c = pd.read_csv(file_c)

# Extract unique tx_hash values from both files
tx_hash_a = set(df_a["tx_hash"])
tx_hash_b = set(df_b["tx_hash"])

# Find tx_hash values that exist in A but not in B
missing_in_b_hashes = tx_hash_b - tx_hash_a

# # Filter rows in A where tx_hash is missing in B
missing_in_b = df_a[df_a["tx_hash"].isin(missing_in_b_hashes)][["tx_hash", "quantity"]]
# sum_missing_in_b = sum(missing_in_b["quantity"].apply(str_to_float))

# # Save to CSV
# missing_in_b.to_csv(f"missing_tx_hash_in_{file_b}.csv", index=False)

# missing_in_b_unique_tx_hash = missing_in_b["tx_hash"].unique()


# Print summary
# print(f"Sum quantity of missing tx_hash in {file_a}: {sum_missing_in_b}")
print(f"Total missing tx_hash in {file_b}: {len(missing_in_b_hashes)}")

print(missing_in_b_hashes)

# print(f"Saved results to missing_in_{file_b}.csv")


# sum_a = sum(df_a["quantity"].apply(str_to_float))
# sum_b = sum(df_b["quantity"].apply(str_to_float))
# sum_c = sum(df_c["quantity"].apply(str_to_float))

# print(f"Sum quantity in {file_a}: {sum_a}")
# print(f"Sum quantity in {file_b}: {sum_b}")
# print(f"Sum quantity in {file_c}: {sum_c}")
