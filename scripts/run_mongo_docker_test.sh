#!/bin/bash

docker run --name mongo -d -p 27011:27011 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \                                                       ─╯
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    mongo