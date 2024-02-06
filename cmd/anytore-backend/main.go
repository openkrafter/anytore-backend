package main

import (
	"github.com/openkrafter/anytore-backend/controller"

	_ "github.com/openkrafter/anytore-backend/config"  // init config
	_ "github.com/openkrafter/anytore-backend/logger"  // init logger
	_ "github.com/openkrafter/anytore-backend/service" // init DynamoDB client
)

func main() {
	controller.Run()
}
