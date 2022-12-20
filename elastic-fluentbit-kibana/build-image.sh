#!/usr/bin/env bash

echo "Building go code..."
# go build -ldflags '-s -w -extldflags "-static"' -o server .
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o server .

echo -e "\nBuilding docker image..."
# docker build -t server:latest . || exit 1

# echo -e "\nStarting docker container...\n"
# docker run --rm -it server:latest

echo -e "\nStarting server with docker-compose...\n"
docker-compose up --force-recreate --build --remove-orphan