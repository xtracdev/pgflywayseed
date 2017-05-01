FROM openjdk:8

ENV HTTP_PROXY ${HTTP_PROXY}
ENV HTTPS_PROXY ${HTTPS_PROXY}
ENV http_proxy ${HTTP_PROXY}
ENV https_proxy ${HTTPS_PROXY}

RUN apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y \
    netcat \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /opt/flyway

RUN wget https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/4.0.3/flyway-commandline-4.0.3-linux-x64.tar.gz && tar xvfz flyway-commandline-4.0.3-linux-x64.tar.gz

ADD postgresql-9.4.1208.jre6.jar /opt/flyway/flyway-4.0.3/drivers/postgresql-9.4.1208.jre6.jar
ADD schema.zip /opt/schema/schema.zip
WORKDIR /opt/schema
RUN unzip ./schema.zip
ADD wait_for_pg.sh /opt/wait_for_pg.sh

ENTRYPOINT ["/opt/wait_for_pg.sh"]
