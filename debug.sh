#!/bin/bash

# Array of Docker containers
declare -a containers=( "oj-judged" "oj-origin" "oj-persistence" "oj-oj" "oj-gateway")

# Array of Docker images
declare -a images=( "oj/judged:local")
#"oj/origin:local" "oj/persistence:local" "oj/oj:local"
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

