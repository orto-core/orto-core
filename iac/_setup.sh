#!/bin/bash

codeDeployAgentInstall() {
	sudo apt update
	sudo apt install -y ruby-full wget

	cd /home/ubuntu
	wget https://aws-codedeploy-us-east-1.s3.us-east-1.amazonaws.com/latest/install -O install
	chmod +x ./install
	sudo ./install auto

	sudo systemctl start codedeploy-agent
	sudo systemctl status codedeploy-agent
}

dockerInstall() {
	# Add Docker's official GPG key:
	sudo apt-get update
	sudo apt-get install -y ca-certificates curl
	sudo install -m 0755 -d /etc/apt/keyrings
	sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
	sudo chmod a+r /etc/apt/keyrings/docker.asc

	# Add the repository to Apt sources:
	echo \
		"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" |
		sudo tee /etc/apt/sources.list.d/docker.list >/dev/null
	sudo apt-get update

	# Install
	sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

	# Add Docker to root
	sudo usermod -aG docker ubuntu

	# Restart Docker service
	sudo systemctl restart docker
}

if systemctl status codedeploy-agent >/dev/null 2>&1; then
	echo "CodeDeploy agent is already running."
else
	echo "Installing CodeDeploy agent..."
	codeDeployAgentInstall
fi

if systemctl status docker >/dev/null 2>&1; then
	echo "Docker Engine is already running."
else
	sleep 30
	echo "Installing Docker Engine..."
	dockerInstall
fi

# Debugging: Check Docker installation status
echo "Docker installation status:"
sudo systemctl status docker
