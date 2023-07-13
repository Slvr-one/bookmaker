#!/bin/bash
set -eux

SERVER=$1
RETRY=$2
PORT=$3

# echo -e "\n-- gonna try $RETRY times: --\n"
# wget --retry-connrefused --waitretry=2 --read-timeout=10 --timeout=10 -t $RETRY "$SERVER:$PORT/home"

curl -X POST -H "Content-Type: application/json" -d '{
    "name": "My Bet",
    "amount": 100,
    "horse_id": 1
}' http://localhost:8080/bets

curl http://localhost:8080/bets


golint my-app
