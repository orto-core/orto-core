#!/bin/bash

useKubernetes=false
network="orto-network"
services=()

if [ "$useKubernetes" = true ]; then
	for service in "${services[@]}"; do
		kubectl delete deployment "$service-deployment"
	done
else
	docker volume prune --force
	docker system prune -a -f

	if ! docker network inspect $network &>/dev/null; then
		docker network create $network
	fi
fi
