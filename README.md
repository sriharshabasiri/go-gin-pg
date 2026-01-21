import requests
import csv
from datetime import datetime, timedelta, timezone
import time

DD_API_KEY = "YOUR_API_KEY"
DD_APP_KEY = "YOUR_APP_KEY"

QUERY = "sum:trace.dd.dynamic.span.hits{env:prod,service:abc}.as_count().rollup(sum,60)"
API_URL = "https://api.datadoghq.com/api/v1/query"

HEADERS = {
    "DD-API-KEY": DD_API_KEY,
    "DD-APPLICATION-KEY": DD_APP_KEY
}

START_DATE = datetime(2025, 7, 1, tzinfo=timezone.utc)
END_DATE   = datetime(2025, 10, 1, tzinfo=timezone.utc)

OUTPUT_FILE = "calls_per_min_jul_to_oct_2025.csv"

def unix_ts(dt):
    return int(dt.timestamp())

with open(OUTPUT_FILE, "w", newline="") as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow(["timestamp", "calls_per_min"])

    current = START_DATE

    while current < END_DATE:
        next_hour = current + timedelta(hours=1)

        params = {
            "from": unix_ts(current),
            "to": unix_ts(next_hour),
            "query": QUERY
        }

        response = requests.get(API_URL, headers=HEADERS, params=params)
        response.raise_for_status()
        data = response.json()

        if "series" in data and data["series"]:
            points = data["series"][0]["pointlist"]

            for point in points:
                ts_ms, value = point
                if value is None:
                    continue

                ts = datetime.fromtimestamp(ts_ms / 1000, tz=timezone.utc)
                formatted_ts = ts.strftime("%m/%d/%Y %H:%M")

                writer.writerow([formatted_ts, int(value)])

        current = next_hour

        # polite pause to avoid rate limits
        time.sleep(0.2)

print(f"CSV written to {OUTPUT_FILE}")
