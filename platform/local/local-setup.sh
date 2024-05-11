#!/bin/sh

docker build -t anytore/backend:latest .
docker-compose -f platform/local/docker-compose-database.yml up -d
sleep 1
docker-compose -f platform/local/docker-compose-backend.yml up -d
sleep 1

platform/local/scripts/user-db-setup.sh
platform/local/scripts/user-db-add-localadmin.sh

platform/local/scripts/dynamodb-local-setup.sh
