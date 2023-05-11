#!/bin/bash
set -eux

#for mongo secret:
# kubectl create secret generic mongodb-credentials \
#   --from-literal=username=dviross \
#   --from-literal=password=secretpass

# helm install my-api-server . \
#     --set mongoDbUrl=mongodb://dviross:secretpass@mongodb-service.default.svc.cluster.local:27017/mydb
name=$1
tag=$2

acc="839821061981"
region="eu-central-1" #frankfurt
repo="$acc.dkr.ecr.$region.amazonaws.com"

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
  --namespace=app
  
# kubectl create secret docker-registry regcred \
#   --docker-server=839821061981.dkr.ecr.eu-central-1.amazonaws.com \
#   --docker-username=AWS \
#   --docker-password=$(aws ecr get-login-password) \
#   --namespace=app
