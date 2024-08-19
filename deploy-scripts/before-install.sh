#!/bin/bash

useKubernetes=false

if [ "$useKubernetes" = true ]; then
	for service in "${services[@]}"; do
		kubectl delete deployment "$service-deployment"
	done
else
	docker volume prune --force && docker system prune -f
fi
