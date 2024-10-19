#!/bin/bash

# delete all containers
docker rm $(docker ps -a -f "status=created" -q)

# prune dangling imagesd
docker rmi $(docker images -f "dangling=true" -q)

# Find and remove backend images
docker images | grep 'moviemate-backend' | awk '{print $3}' | xargs docker rmi -f