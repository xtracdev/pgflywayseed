This project provides a method to seed a postgres schema using
flyway as part of a circle ci build.

To use, zip up your flyway project into a file named schema.zip, and 
place it in this directory prior to doing your docker build. The docker
build will unzip it into /opt/schema, and the script that invokes flyway
assumes the migration directory is in /opt/schema/db.

When running in a Circle CI 2.0 build job, if postgres is run as a container
in the build, it will be available from this container on localhost.

You will need to provide the DB_HOST, DB_PORT, USER, PASSWORD, and URL
environment variables. DB_HOST and DB_PORT are used as inputs to nc
for determining when postgres is available, and the rest are used as 
flyway inputs.

Below is a snipped from flyway config showing how it fits into a build.

<pre>
jobs:
  build:
    working_directory: /go/src/github.com/xtracdev/pgeventstore
    docker:
      - image: golang:1.8.1-onbuild
      - image: postgres:9.6.2
        environment:
          POSTGRES_PASSWORD: password
      - image: xtracdev/pgflywayseed
        environment:
          DB_HOST: localhost
          DB_PORT: 5432
          USER: postgres
          PASSWORD: password
          URL: jdbc:postgresql://localhost:5432/postgres
</pre>

