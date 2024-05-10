package dynamodb

import "github.com/openkrafter/anytore-backend/logger"

var DClient DynamoDBService

func InitDynamoDbClient() {
	var err error
	DClient, err = NewDynamoDBClient()
	if err != nil {
		errMsg := "Failed to init DynamoDBClient."
		logger.Logger.Error(errMsg, logger.ErrAttr(err))
		panic(errMsg)
	}
}
