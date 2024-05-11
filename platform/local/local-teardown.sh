#!/bin/sh

platform/local/scripts/dynamodb-local-teardown.sh

platform/local/scripts/user-db-teardown.sh

sleep 2

docker-compose -f platform/local/docker-compose.yml down
