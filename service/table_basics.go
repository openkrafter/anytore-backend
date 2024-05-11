package service

import (
	anytoreDynamodb "github.com/openkrafter/anytore-backend/database/dynamodb"
)

type TableBasics struct {
	DynamoDbClient anytoreDynamodb.DynamoDBService
	TableName      string
}

func NewTableBasics(tableName string) (*TableBasics, error) {
	basics := new(TableBasics)
	basics.DynamoDbClient = anytoreDynamodb.DClient
	basics.TableName = tableName

	return basics, nil
}
