package controller

import (
	"os"
	"testing"

	anytoreConfig "github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/logger"
	testenvironment "github.com/openkrafter/anytore-backend/test/environment"
)

func setup() error {
	logger.InitLogger()
	logger.Logger.Info("controller package test setup")

	anytoreConfig.InitConfig()

	err := testenvironment.SetupDynamoDbClient()
	if err != nil {
		logger.Logger.Error("SetupDynamoDbClient failed.", logger.ErrAttr(err))
		return err
	}

	return nil
}

func teardown() error {
	logger.Logger.Info("controller package test done")
	return nil
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		logger.Logger.Error("Failed to setup.", logger.ErrAttr(err))
	}

	ret := m.Run()

	if err := teardown(); err != nil {
		logger.Logger.Error("Failed to teardown.", logger.ErrAttr(err))
	}

	os.Exit(ret)
}
