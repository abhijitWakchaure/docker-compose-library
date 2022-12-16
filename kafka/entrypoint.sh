#!/usr/bin/env bash

# Start the ZooKeeper service
# Note: Soon, ZooKeeper will no longer be required by Apache Kafka.
/${KAFKA_SRC_DIR}/bin/zookeeper-server-start.sh /${KAFKA_SRC_DIR}/config/zookeeper.properties &

/post-startup.sh &

# Start the Kafka broker service
/${KAFKA_SRC_DIR}/bin/kafka-server-start.sh /${KAFKA_SRC_DIR}/config/server.properties

