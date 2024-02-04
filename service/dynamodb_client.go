package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	anytoreConfig "github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/logger"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewTableBasics(tableName string) (*TableBasics, error) {
	basics := new(TableBasics)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(anytoreConfig.Config.AWS_REGION),
	)
	if err != nil {
		logger.Logger.Error("Load aws config error.", logger.ErrAttr(err))
		return nil, err
	}

	basics.DynamoDbClient = dynamodb.NewFromConfig(cfg)
	basics.TableName = tableName

	return basics, nil
}