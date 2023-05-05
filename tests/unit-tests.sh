#!/bin/bash
set -eux

SERVER=$1
RETRY=$2
PORT=$3

echo -e "\n-- gonna try $RETRY times: --\n"
wget --retry-connrefused --waitretry=2 --read-timeout=10 --timeout=10 -t $RETRY "$SERVER:$PORT/home"
