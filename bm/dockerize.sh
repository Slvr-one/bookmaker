#!/bin/bash
set -eux

# bilding the app again & reinstalling dependencies

rm -rf go.mod go.sum
go mod init src/main.go

#install known dependencies
pushd src
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/mongo/readpref
popd

#build module & image
go build src/main.go && echo -e "\n ----go build success!!---- \n" || echo -e "\n ----go build failed!!---- \n"
docker build -t $1 . ##--build-arg serverPort=5555

#docker run -d --name bookmaker -p 5000:5000 bookmaker
# docker rm -f bookMaker
