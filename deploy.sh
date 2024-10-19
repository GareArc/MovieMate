#!/bin/bash

docker compose pull

docker rm $(docker ps -a -f "status=created" -q)

docker compose up -d
