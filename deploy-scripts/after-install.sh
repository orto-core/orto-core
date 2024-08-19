#!/bin/bash

services=()

useKubernetes=false

if [ "$useKubernetes" = true ]; then
	switchDir "/home/projects/orto-micro-service"
	for service in "${services[@]}"; do
		kubectl apply -f "./kubernetes/deployments/$service.yaml"
	done
else
	cd "/home/projects/orto-micro-service"
	sudo docker compose -f "docker-compose.yml" up --remove-orphans -d
fi
