#!/bin/bash
set -eux

EC2_NAME=$1
INSTANCE_ID=$(aws ec2 describe-instances \
    --filters "Name=tag:Name,Values=$EC2_NAME"  \
    "Name=instance-state-name,Values=running" | \
    jq -r ".Reservations[] | .Instances[] | .InstanceId")

#get instance start time - formated--
INSTANCE_START_TIME=$(aws ec2 describe-instances \
    --filter Name=instance-state-name,Values=running \
    --output table \
    --query "Reservations[].Instances[].{Name: Tags[?Key == 'Name'].Value | \
    [0], Id: InstanceId, State: State.Name, Type: InstanceType, Start: LaunchTime}" | \
    grep "$EC2_NAME" | cut -d "|" -f4 | cut -d "T" -f 2 | cut -d "." -f1 | cut -d "+" -f1)

IFS=: read h m s <<<"$INSTANCE_START_TIME"
SECONDS_THEN=$((10#$s+10#$m*60+10#$h*3600))

NOW=$(date +"%T")
IFS=: read h m s <<<"$NOW"
SECONDS_NOW=$((10#$s+10#$m*60+10#$h*3600))

DELTA_SECONDS=$(($SECONDS_NOW - $SECONDS_THEN))
DELTA_MINUTES=$((($DELTA_SECONDS / 60) - 1))

if [ $DELTA_MINUTES -ge 15 ] || [ $DELTA_MINUTES -le -15 ]; then
    echo "$INSTANCE_ID is older than 15 minutes - delete -"
    aws ec2 terminate-instances --instance-ids $INSTANCE_ID
    echo "$INSTANCE_ID - deleted -"
else
    echo "not gonna delete instance"
fi

echo  "now = $NOW"
echo "instance start time = $INSTANCE_START_TIME"

##########################################################################################

# aws ec2 describe-instances \
#     --filter Name=instance-state-name,Values=running \
#     --output table \
#     --query "Reservations[].Instances[].{Name: Tags[?Key == 'Name'].Value | \
#     [0], Id: InstanceId, State: State.Name, Type: InstanceType, Start: LaunchTime}" 

