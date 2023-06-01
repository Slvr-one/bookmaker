#!/bin/bash
set -eux

# deploy app for with nginx and mongodb containers:
docker compose down && echo -e "\n ###down### \n"

docker compose up --build -d && echo -e "\n ###up### \n"
