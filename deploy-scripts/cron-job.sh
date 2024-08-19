#!/bin/bash

SCRIPT_PATH="./backup-logs.sh"

CRON_JOB="0 0 * * * $SCRIPT_PATH"

(crontab -l | grep -F "$SCRIPT_PATH") || (
	(
		crontab -l 2>/dev/null
		echo "$CRON_JOB"
	) | crontab -
	echo "Cron job added: $CRON_JOB"
)
