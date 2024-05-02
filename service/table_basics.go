package service

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	anytoreDynamodb "github.com/openkrafter/anytore-backend/database/dynamodb"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewTableBasics(tableName string) (*TableBasics, error) {
	basics := new(TableBasics)
	basics.DynamoDbClient = anytoreDynamodb.DynamoDbClient
	basics.TableName = tableName

	return basics, nil
}
