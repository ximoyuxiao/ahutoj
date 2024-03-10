#!/bin/bash

# Array of Docker containers
declare -a containers=( "oj-judged" "oj-origin"  "oj-oj" "oj-gateway")

# Array of Docker images
declare -a images=( "oj/judged:local" "oj/origin:local" "oj/oj:local:local" "oj-gateway:local")
# Stop Docker containers
for container in "${containers[@]}"
do
   echo "Stopping $container"
   sudo docker stop $container
done

# Remove Docker images
for image in "${images[@]}"
do
   echo "Removing $image"
   sudo docker rmi -f $image
done

sudo docker-compose up -d

echo "Script completed"

