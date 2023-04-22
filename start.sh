#!/bin/sh

set -e

echo "run db migration"

aws secretsmanager get-secret-value --secret-id recipe --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > /app/app.env

source /app/app.env

/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"