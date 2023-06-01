#!/bin/bash
set -euxo pipefail

#run after init, to exec from containers
jenpass=$(docker compose exec jenkins \
cat /var/jenkins_home/secrets/initialAdminPassword)
labpass=$(docker compose exec gitlab \
cat etc/gitlab/initial_root_password | grep "Password:" | awk '{print $2}')

echo -e "\n jen-pass is $jenpass \n"
echo -e "\n lab-pass is $labpass \n"


