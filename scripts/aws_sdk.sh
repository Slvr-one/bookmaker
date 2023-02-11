#!/bin/bash
set -eu

go get github.com/aws/aws-sdk-go-v2\

modules=("github.com/aws/aws-sdk-go-v2/config"\
    "github.com/aws/aws-sdk-go-v2/aws"\
    "github.com/aws/aws-sdk-go-v2/service/dynamodb")

for i in "${modules[@]}"
    do
#    : 
    go get $i
done