#!/bin/bash
set -eu

# curl -sSL <raw_script_url> | bash

# wget -qO- <raw_script_url> | bash

go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv

export MONGODB_URI='<your atlas connection string>'


pushd ~/_worklplace/bookmaker
docker compose up --build -d

popd