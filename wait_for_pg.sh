#!/bin/bash

set -e

/opt/connectdb


>&2 echo "Postgres up - executing command"
exec /opt/flyway/flyway-4.0.3/flyway -user=$DB_USER -password=$DB_PASSWORD -locations=filesystem:/opt/schema/db -driver=org.postgresql.Driver -url=$DB_URL migrate
