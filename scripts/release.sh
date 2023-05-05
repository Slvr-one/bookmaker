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

region="eu-central-1" #frankfurt
repo="514095112279.dkr.ecr.$region.amazonaws.com"

#install awscli(.sh)

aws ecr get-login-password \
    --region $region | \
    docker login --username AWS \
    --password-stdin $repo

docker build -t $name .
docker tag $name:latest $repo/$name:$tag
docker push $repo/$name:$tag