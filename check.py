import pandas as pd # type: ignore

# File paths
file_b = "mainchain-weth-withdrawals.csv"
file_a = "ronin-weth-request-withdrawals.csv"

# Load CSV files into DataFrames
df_a = pd.read_csv(file_a)
df_b = pd.read_csv(file_b)

# Extract unique receipt_hash values from both files
receipt_hash_a = set(df_a["receipt_hash"])
receipt_hash_b = set(df_b["receipt_hash"])

# Find receipt_hash values that exist in A but not in B
missing_in_b_hashes = receipt_hash_a - receipt_hash_b

# Filter rows in A where receipt_hash is missing in B
missing_in_b = df_a[df_a["receipt_hash"].isin(missing_in_b_hashes)][
    ["receipt_hash", "tx_hash", "receipt_id"]
]

# Save to CSV
missing_in_b.to_csv("missing_in_b.csv", index=False)

# Print summary
print(f"Total missing receipt_hash in {file_b}: {len(missing_in_b_hashes)}")
print(f"Saved results to missing_in_{file_b}.csv")
