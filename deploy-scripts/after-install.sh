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
	docker compose up -f "docker-compose.yml"
fi
