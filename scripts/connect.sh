#!/bin/bash
set -eu

go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv

export MONGODB_URI='<your atlas connection string>'
