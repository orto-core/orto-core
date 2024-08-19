#!/bin/bash

LOG_DIR="/var/log"
S3_BUCKET="s3://orto-deployment-assets"
BACKUP_DIR="/home/nathan/logs"
RETENTION_DAYS=2

backupLogs() {
	mkdir -p $BACKUP_DIR

	find $LOG_DIR -type f -mtime +$RETENTION_DAYS -exec mv {} $BACKUP_DIR \;

	aws s3 sync $BACKUP_DIR $S3_BUCKET

	find $LOG_DIR -type f -mtime +$RETENTION_DAYS -exec rm {} \;

	rm -rf $BACKUP_DIR
}

backupLogs
