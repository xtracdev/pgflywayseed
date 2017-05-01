#!/bin/bash

set -e

host="$DB_HOST"
port="$DB_PORT"

ls -R /opt/schema

echo DB_HOST is $host
echo DB_PORT=$port

until nc -z $host $port; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres up - executing command"
exec /opt/flyway/flyway-4.0.3/flyway -user=$USER -password=$PASSWORD -locations=filesystem:/opt/schema/db -driver=org.postgresql.Driver -url=$URL migrate
