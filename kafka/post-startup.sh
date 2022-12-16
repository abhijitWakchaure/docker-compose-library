#!/usr/bin/env bash

echo "Sleeping for 10s and waiting for Kafka server to start..."
sleep 10
/${KAFKA_SRC_DIR}/bin/kafka-topics.sh --create --topic test --bootstrap-server localhost:9092