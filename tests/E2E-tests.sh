#!/bin/bash
set -eux

SERVER=$1
RETRY=$2
PORT=$3

echo -e "ping...\n"
wget --retry-connrefused \
--waitretry=2 --read-timeout=10 \
--timeout=10 -t $RETRY "$SERVER:$PORT"/health

# health=$(cat health)
# if "health" in $health; then
#     echo "got health"
# else exit 1
# fi
wget --retry-connrefused --waitretry=2 --read-timeout=10 --timeout=10 -t $RETRY "$SERVER:$PORT/horses"

horses=$(cat horses) \
    && echo -e "\n-------- found $SERVER server on port $PORT - $horses horses------\n" \
    && rm -rf horses || echo "found an alien"