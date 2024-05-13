#!/bin/sh

aws dynamodb delete-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItem

aws dynamodb delete-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItemCounter

aws dynamodb delete-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingRecord

aws dynamodb delete-table \
  --no-cli-pager \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingRecordCounter
