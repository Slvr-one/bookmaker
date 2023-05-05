#!/bin/bash
set -euo pipelinefail
# sudo -i
# export DEBIAN_FRONTEND=noninteractive

sudo apt update -qqy

sudo apt-get install -qqy \
    apt-transport-https git \
    ca-certificates curl gnupg2 \
    software-properties-common \
    bzip2 sudo wget vim tree  \
    lsb-release grub-efi-amd64-bin

mkdir -p /etc/apt/keyrings
curl -fsSL \
    https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg > /tmp/dkey
    sudo apt-key add -q /tmp/dkey

# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
# echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
#   $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
    $(lsb_release -cs) \
    stable"

# sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"


apt-cache policy docker-ce && sudo apt install docker-ce
# sudo apt-get update  -qqy \
#   && sudo apt-get install -qqy docker-ce docker-ce-cli containerd.io docker-compose-plugin

sudo usermod -aG docker $USER
newgrp docker

#su - $USER

sudo service docker start
sudo systemctl status docker


# echo "Install Jenkins"
# sudo wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat-stable/jenkins.repo
# sudo rpm --import https://jenkins-ci.org/redhat/jenkins-ci.org.key
# sudo yum install -y jenkins
# sudo usermod -a -G docker jenkins
# sudo chkconfig jenkins on
# sudo service jenkins start

# echo "Install Terraform"
# wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
# echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com impish main" | \
#     sudo tee /etc/apt/sources.list.d/hashicorp.list
# sudo apt update && sudo apt install terraform

# echo "Install AWScli"
# curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
# unzip awscliv2.zip
# sudo ./aws/install

# sudo apt install awscli -y

# ****** #
# sudo yum -y update

# echo "Install Java JDK 8"
# sudo yum remove -y java
# sudo yum install -y java-1.8.0-openjdk

# echo "Install Maven"
# sudo yum install -y maven 


# echo "Install Docker engine"
# sudo yum update -y
# sudo yum install docker -y
# sudo sudo chkconfig docker on