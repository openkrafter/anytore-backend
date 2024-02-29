package main

import (
	"github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/controller"
	"github.com/openkrafter/anytore-backend/logger"

	"github.com/openkrafter/anytore-backend/dynamodb" // init DynamoDB client
)

func main() {
	logger.InitLogger()
	config.InitConfig()
	dynamodb.InitDynamoDbClient()

	controller.Run()
}
