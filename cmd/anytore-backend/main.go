package main

import (
	"github.com/openkrafter/anytore-backend/auth"
	"github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/controller"
	"github.com/openkrafter/anytore-backend/logger"

	"github.com/openkrafter/anytore-backend/database/dynamodb"
	"github.com/openkrafter/anytore-backend/database/sqldb"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.InitLogger()
	config.InitConfig()
	dynamodb.InitDynamoDbClient()

	err := sqldb.InitSQLDBClient()
	if err != nil {
		logger.Logger.Error("Failed to initialize SQLDB client", logger.ErrAttr(err))
		return
	}

	auth.InitPassHasher()

	controller.Run()
}
