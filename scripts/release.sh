#!/bin/bash
set -euxo pipefail

#for mongo secret:
# kubectl create secret generic mongodb-credentials \
#   --from-literal=username=dviross \
#   --from-literal=password=secretpass

# helm install my-api-server . \
#     --set mongoDbUrl=mongodb://dviross:secretpass@mongodb-service.default.svc.cluster.local:27017/mydb

export name=$1
export tag=$2

export work="./"
# export repo="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
export lab="./lab_infra" # "~/_workplace/portfolio/infra/lab/" 

export acc="839821061981"
export region="eu-central-1" #frankfurt
export repo="$acc.dkr.ecr.$region.amazonaws.com"

#install awscli(.sh)

aws ecr get-login-password \
    --region $region | \
    docker login --username AWS \
    --password-stdin $repo

docker build -t $name .
docker tag $name:latest $repo/$name:$tag
docker push $repo/$name:$tag

kubectl create secret docker-registry regcred \
  --docker-server=${acc}.dkr.ecr.${region}.amazonaws.com \
  --docker-username=AWS \
  --docker-password=$(aws ecr get-login-password) \
  --namespace=app || true # --dry-run=client -o yaml | kubectl create -f - \
  # --ignore-not-found=true

./scripts/dc_up.sh
./scripts/get_pass.sh


#deploy testing lab (Jenkins server, Gitlab server, Artifactory and maven development environment):
pushd $lab

./../scripts/dc_up.sh

popd


# kubectl create secret docker-registry regcred \
#   --docker-server=839821061981.dkr.ecr.eu-central-1.amazonaws.com \
#   --docker-username=AWS \
#   --docker-password=$(aws ecr get-login-password) \
#   --namespace=app
