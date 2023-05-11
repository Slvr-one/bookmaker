#!/bin/bash
set -eux
# please provide a tag parameter for this script

# go install github.com/cosmtrek/air@latest


#build module & image
go build src/main.go && echo -e "\n ----go build success!!---- \n" || echo -e "\n ----go build failed!!---- \n"
docker build -t bm:$1 . ##--build-arg serverPort=5555

docker run -d --name bookmaker -p 2000:5000 bm:$1
# docker rm -f bookmaker
