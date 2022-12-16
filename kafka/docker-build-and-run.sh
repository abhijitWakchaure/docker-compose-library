#!/usr/bin/env bash

KAFKA_TGZ=kafka_2.13-3.2.1.tgz
KAFKA_SRC_DIR=${KAFKA_TGZ%.*}


docker-compose build --build-arg KAFKA_TGZ=${KAFKA_TGZ} --build-arg KAFKA_SRC_DIR=${KAFKA_SRC_DIR}

docker-compose up --force-recreate