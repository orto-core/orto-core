#!/bin/bash

useKerbernetes=false
services=()

if [ "$useKerbernetes" = true ]; then
	for service in "${services[@]}"; do
		kubectl rollout status deployment "$service-deployment"
	done
else
	docker ps
fi
