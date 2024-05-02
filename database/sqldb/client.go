package sqldb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/openkrafter/anytore-backend/logger"
)

var SQLDBClient SQLDBService

func InitSQLDBClient() error {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/anytore?parseTime=true", dbUser, dbPassword, dbHost)
	sqldb, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Logger.Error("Connect Error: MySQL.", logger.ErrAttr(err))
		panic("Failed to start anytore.")
		// return err
	}

	logger.Logger.Debug("Connected to MySQL database")

	SQLDBClient = NewSQLDB(sqldb)

	return nil
}
