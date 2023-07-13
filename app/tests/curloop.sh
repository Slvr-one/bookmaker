#!/usr/bin/env bash
set -eu

export APP="$1"
echo "curling URL $APP in a loop..."

while true
do 
    curl $APP 
    sleep 10
done