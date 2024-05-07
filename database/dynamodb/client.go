package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/openkrafter/anytore-backend/logger"
)

var DynamoDbClient *dynamodb.Client

func InitDynamoDbClient() {
	logger.Logger.Debug("Init DynamoDB client.")

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Logger.Error("Load aws config error.", logger.ErrAttr(err))
		panic("Failed to start anytore.")
	}

	DynamoDbClient = dynamodb.NewFromConfig(cfg)
}
