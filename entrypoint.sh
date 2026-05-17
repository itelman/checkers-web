#!/bin/bash

# Load environment variables from .env file
# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs)

docker build -t checkers-api .
docker run -d --name checkers-api -p 8888:8888 checkers-api