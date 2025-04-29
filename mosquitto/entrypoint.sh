#!/usr/bin/env bash
set -e

# Generate passwd file only on first run
PASSWD_FILE=/mosquitto/config/passwd
if [ ! -f "$PASSWD_FILE" ]; then
  /usr/local/bin/generate_mosquitto_passwd.sh
fi

# Hand off to original entrypoint
exec /docker-entrypoint.sh "$@"
