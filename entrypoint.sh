#!/bin/bash

# Load environment variables from .env file
# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs)

docker build -t checkers .
docker run -d --name checkers -p 3000:3000 checkers