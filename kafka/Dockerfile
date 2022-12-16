FROM amazoncorretto:8

# RUN yum update -y && yum install tar -y

WORKDIR /

# Filename for kafka .tgz file downloaded from https://kafka.apache.org/quickstart
ARG KAFKA_TGZ=${KAFKA_TGZ}
ARG KAFKA_SRC_DIR=${KAFKA_SRC_DIR}

ENV KAFKA_TGZ ${KAFKA_TGZ}
ENV KAFKA_SRC_DIR ${KAFKA_SRC_DIR}

ADD ./${KAFKA_TGZ} .
ADD ./entrypoint.sh .
ADD ./post-startup.sh .

ENTRYPOINT [ "/entrypoint.sh" ]