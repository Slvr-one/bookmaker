#!/bin/bash
set -eux

# please provide a tag parameter for this script

# bilding the app & reinstalling dependencies

rm -rf go.mod go.sum
go mod init src/main.go

#install known dependencies
# go install github.com/cosmtrek/air@latest

# pushd src
# go get github.com/gorilla/mux
# go get github.com/joho/godotenv
# go get go.mongodb.org/mongo-driver/bson
# go get go.mongodb.org/mongo-driver/mongo
# go get go.mongodb.org/mongo-driver/mongo/options
# go get go.mongodb.org/mongo-driver/mongo/readpref
# popd

#build module & image
go build src/main.go && echo -e "\n ----go build success!!---- \n" || echo -e "\n ----go build failed!!---- \n"
docker build -t bm:$1 . ##--build-arg serverPort=5555

docker run -d --name bookmaker -p 2000:5000 bm:$1
# docker rm -f bookmaker
