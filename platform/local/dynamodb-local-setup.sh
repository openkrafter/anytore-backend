#!/bin/sh

aws dynamodb create-table \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItem \
  --attribute-definitions AttributeName=Id,AttributeType=N \
  --key-schema AttributeName=Id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --table-class "STANDARD" \
  --no-deletion-protection-enabled

aws dynamodb create-table \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItemCounter \
  --attribute-definitions AttributeName=CountKey,AttributeType=S \
  --key-schema AttributeName=CountKey,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --table-class "STANDARD" \
  --no-deletion-protection-enabled

aws dynamodb put-item \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItemCounter \
  --item '{ "CountKey": {"S": "key"}, "CountNumber": {"N": "0"} }'
