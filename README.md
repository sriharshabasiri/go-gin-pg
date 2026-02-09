import os
import shutil
import sys
from pathlib import Path
import hashlib
import fcntl

VERSION_FILE = Path("flyway_version.txt")
LOCK_FILE = Path("flyway_version.lock")

INCOMING_DIR = Path("incoming_sql")
FLYWAY_DIR = Path("flyway_sql")
ARCHIVE_DIR = Path("archive_sql")


def atomic_write(path, content):
    tmp = path.with_suffix(".tmp")
    with open(tmp, "w") as f:
        f.write(content)
        f.flush()
        os.fsync(f.fileno())
    os.replace(tmp, path)


def get_current_version():
    if not VERSION_FILE.exists():
        return 0
    return int(VERSION_FILE.read_text().strip())


def calculate_checksum(file_path):
    h = hashlib.sha256()
    with open(file_path, "rb") as f:
        for chunk in iter(lambda: f.read(8192), b""):
            h.update(chunk)
    return h.hexdigest()


def main():
    if not INCOMING_DIR.exists():
        print("No incoming SQL directory. Exiting.")
        return

    sql_files = sorted(f for f in INCOMING_DIR.iterdir() if f.suffix == ".sql")
    if not sql_files:
        print("No SQL files to process.")
        return

    LOCK_FILE.touch(exist_ok=True)

    with open(LOCK_FILE, "r+") as lock_fd:
        fcntl.flock(lock_fd, fcntl.LOCK_EX)

        current_version = get_current_version()
        next_version = current_version + 1

        print(f"Current file-based Flyway version: {current_version}")

        FLYWAY_DIR.mkdir(parents=True, exist_ok=True)
        ARCHIVE_DIR.mkdir(parents=True, exist_ok=True)

        for sql_file in sql_files:
            new_name = f"V{next_version}__{sql_file.stem}.sql"
            shutil.copy2(sql_file, FLYWAY_DIR / new_name)
            shutil.move(sql_file, ARCHIVE_DIR / sql_file.name)
            print(f"Generated: {new_name}")
            next_version += 1

        atomic_write(VERSION_FILE, str(next_version - 1))

        fcntl.flock(lock_fd, fcntl.LOCK_UN)

    print("File-based migration preparation completed.")


if __name__ == "__main__":
    main()
