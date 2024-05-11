#!/bin/sh

aws dynamodb create-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItem \
  --attribute-definitions AttributeName=Id,AttributeType=N AttributeName=UserId,AttributeType=N \
  --key-schema AttributeName=Id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --table-class "STANDARD" \
  --no-deletion-protection-enabled \
  --global-secondary-indexes \
    '[
      {
        "IndexName": "UserIdIndex",
        "KeySchema": [
          {"AttributeName": "UserId", "KeyType": "HASH"}
        ],
        "Projection": {
          "ProjectionType": "ALL"
        },
        "ProvisionedThroughput": {
          "ReadCapacityUnits": 1,
          "WriteCapacityUnits": 1
        }
      }
    ]'

aws dynamodb create-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItemCounter \
  --attribute-definitions AttributeName=CountKey,AttributeType=S \
  --key-schema AttributeName=CountKey,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --table-class "STANDARD" \
  --no-deletion-protection-enabled
