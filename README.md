import pandas as pd
import numpy as np

# -----------------------------
# CONFIGURATION
# -----------------------------
CSV_PATH = "txn_data.csv"

ACTIVE_START_HOUR = 6
ACTIVE_END_HOUR = 23

MIN_TXN_THRESHOLD = 10          # single digit or zero is abnormal
DEVIATION_THRESHOLD = -0.30     # 30% drop

# -----------------------------
# LOAD DATA
# -----------------------------
df = pd.read_csv(CSV_PATH)

# Ensure proper ordering
df = df.sort_values(by=["date", "minute_of_day"]).reset_index(drop=True)

# -----------------------------
# EXPECTED TRANSACTION BASELINE
# -----------------------------
df['expected_txn'] = (
    df.groupby(['minute_of_day', 'is_weekend'])['txn_count']
      .transform('mean')
)

# Avoid divide-by-zero
df['expected_txn'] = df['expected_txn'].replace(0, np.nan)

# -----------------------------
# DEVIATION FROM EXPECTED
# -----------------------------
df['pct_deviation'] = (
    (df['txn_count'] - df['expected_txn']) / df['expected_txn']
)

# -----------------------------
# ROLLING FEATURES (TREND AWARENESS)
# -----------------------------
df['rolling_15'] = df['txn_count'].rolling(window=15, min_periods=1).mean()
df['rolling_60'] = df['txn_count'].rolling(window=60, min_periods=1).mean()

df['pct_from_rolling_60'] = (
    (df['txn_count'] - df['rolling_60']) / df['rolling_60']
)

# -----------------------------
# ACTIVE HOURS FLAG
# -----------------------------
df['is_active_hour'] = (
    (df['hour'] >= ACTIVE_START_HOUR) &
    (df['hour'] <= ACTIVE_END_HOUR)
)

# -----------------------------
# TXN STATUS LABEL
# 0 = NORMAL, 1 = ABNORMAL
# -----------------------------
df['txn_status'] = 0

df.loc[
    (
        df['is_active_hour'] &
        (
            (df['txn_count'] < MIN_TXN_THRESHOLD) |
            (df['pct_deviation'] < DEVIATION_THRESHOLD)
        )
    ),
    'txn_status'
] = 1

# -----------------------------
# CLEANUP FOR ML
# -----------------------------
df.replace([np.inf, -np.inf], np.nan, inplace=True)

df['pct_deviation'].fillna(0, inplace=True)
df['pct_from_rolling_60'].fillna(0, inplace=True)

# -----------------------------
# FINAL FEATURE SET
# -----------------------------
FEATURE_COLUMNS = [
    'txn_count',
    'minute_of_day',
    'is_weekend',
    'expected_txn',
    'pct_deviation',
    'rolling_15',
    'pct_from_rolling_60'
]

X = df[FEATURE_COLUMNS]
y = df['txn_status']

# -----------------------------
# OPTIONAL: SAVE OUTPUT
# -----------------------------
df.to_csv("txn_features_with_labels.csv", index=False)

print("Feature engineering completed.")
print("Abnormal txn ratio:")
print(y.value_counts(normalize=True))
