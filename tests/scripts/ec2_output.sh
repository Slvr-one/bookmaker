#!/bin/bash

#output in json
user_info=`{ aws sts get-caller-identity & aws iam list-account-aliases; } | jq -s ".|add"`
region=`aws ec2 describe-availability-zones --output text --query 'AvailabilityZones[0].[RegionName]'`
services_in_use=`aws ce get-cost-and-usage \
    --time-period Start=$(date "+%Y-%m-01" -d "-1 Month"),End=$(date \
    --date="$(date +'%Y-%m-01') - 1 second" -I) \
    --granularity MONTHLY --metrics UsageQuantity \
    --group-by Type=DIMENSION,Key=SERVICE | jq '.ResultsByTime[].Groups[] | \
    select(.Metrics.UsageQuantity.Amount > 0) | .Keys[0]'`

instances_state_and_type=`aws ec2 describe-instances | \
    jq -r "[[.Reservations[].Instances[]|{ state: .State.Name, type: .InstanceType }]|\
    group_by(.state)|.[]|{state: .[0].state, types: [.[].type]|\
    [group_by(.)|.[]|{type: .[0], count: ([.[]]|length)}] }]"`

cost_per_service=`aws ce get-cost-and-usage \
    --time-period Start=$(date "+%Y-%m-01"),End=$(date --date="$(date +'%Y-%m-01') + 1 month  - 1 second" -I) \
    --granularity MONTHLY --metrics USAGE_QUANTITY BLENDED_COST  \
    --group-by Type=DIMENSION,Key=SERVICE | jq '[ .ResultsByTime[].Groups[] | \
    select(.Metrics.BlendedCost.Amount > "0") | \
    { (.Keys[0]): .Metrics.BlendedCost } ] | sort_by(.Amount) | add'`
