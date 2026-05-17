#!/bin/bash

# Load environment variables from .env file
# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs)

docker build -t checkers .
sleep 5

docker run -d --name checkers -p 8888:8888 checkers