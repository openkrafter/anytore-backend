package main

import (
	"github.com/openkrafter/anytore-backend/controller"

	_ "github.com/openkrafter/anytore-backend/config"   // init config
	_ "github.com/openkrafter/anytore-backend/dynamodb" // init DynamoDB client
	_ "github.com/openkrafter/anytore-backend/logger"   // init logger
)

func main() {
	controller.Run()
}
