#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

# Run the migrations
echo "Running migrations..."
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

# Run the main application
echo "Running the main application..."
exec "$@"