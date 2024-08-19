#!/bin/bash

sudo cp ./nginx.conf /etc/nginx/sites-available/orto-server

sudo ln -s /etc/nginx/sites-available/orto-server /etc/nginx/sites-enabled

sudo nginx -t

if [ $? -eq 0 ]; then
	echo "Added orto nginx reverse proxy configuration file"
	sudo systemctl reload nginx
else
	echo "Nginx configuration for orto server failed..."
fi
