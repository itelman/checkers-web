#!/bin/bash

# Load environment variables from .env file
# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs)

# Start Docker containers in detached mode
docker-compose up --no-start

# Run database migrations
docker-compose start postgres
sleep 5
migrate -database "postgresql://${PG_USER}:${PG_PASS}@localhost:${PG_PORT}/${PG_DB}?sslmode=disable" -path ./migrations/postgres up
sleep 5

docker-compose start backend
