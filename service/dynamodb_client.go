package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewTableBasics(tableName string) (*TableBasics, error) {
	basics := new(TableBasics)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		// log.Fatalf("unable to load SDK config, %v", err)
		fmt.Println("error test")
		return nil, err
	}

	basics.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	basics.TableName = tableName

	return basics, nil
}