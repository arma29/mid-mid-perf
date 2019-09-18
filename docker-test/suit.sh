#!/bin/bash

# chmod +x suit.sh
# Create directory structure
grpc='grpc_client'
rab='rabbit_client'

mkdir -p /home/$USER/Output/$grpc
mkdir -p /home/$USER/Output/$rab

# Remove files before starts
(find ~/Output/$grpc -type f -delete)
(find ~/Output/$rab -type f -delete)

# Update/Create docker image before starts
(docker build --rm --tag=perf-image:v2 .)

#GRPC x30
current='grpc/docker-compose.yml'
for i in {1..30}
do
    (docker-compose -f $current up -d)
    (sleep 10s) #wait writes everything in file
    (docker-compose -f $current down)
done
echo "GRPC Done"

#RABBIT x30
current='rabbitmq/docker-compose.yml'
for i in {1..30}
do
    (docker-compose -f $current up -d)
    (sleep 10s) #wait writes everything in file
    (docker-compose -f $current down)
done
echo "RABBITMQ Done"