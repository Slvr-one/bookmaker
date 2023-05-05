#!/bin/bash
set -eu
. .env

echo "may take a while.."

pushd ~/Downloads/aws-keys/
scp -ri $ec2_ssh_private_key ~/Desktop/ci/scripts/* ubuntu@$ec2_ip:$onubuntu && echo -e "\n copied to ec2, in folder $onubuntu \n" || echo -e "\n ---could not copy to ec2--- \n"
popd
