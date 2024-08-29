#!/bin/bash

DOCKER_LOG_DIR="/var/lib/docker/containers"
S3_BUCKET="s3://orto-deployment-assets"
BACKUP_DIR="/home/projects/logs"
RETENTION_DAYS=2

mkdir -p "$BACKUP_DIR"

backupLogs() {
	for log_dir in "$DOCKER_LOG_DIR"/*; do
		if [ -d "$log_dir" ]; then
			find "$log_dir" -type f -name "*.log" -mtime +$RETENTION_DAYS -exec mv {} "$BACKUP_DIR" \;
		fi
	done

	aws s3 sync "$BACKUP_DIR" "$S3_BUCKET"

	for log_dir in "$DOCKER_LOG_DIR"/*; do
		if [ -d "$log_dir" ]; then
			find "$log_dir" -type f -name "*.log" -mtime +$RETENTION_DAYS -exec rm {} \;
		fi
	done

	rm -rf "$BACKUP_DIR"/*
}

backupLogs
