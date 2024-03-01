#!/bin/sh

aws dynamodb delete-table \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItem

aws dynamodb delete-table \
  --endpoint-url http://localhost:8000 \
  --table-name TrainingItemCounter
