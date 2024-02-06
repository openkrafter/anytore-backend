package config

import (
	"fmt"
	"os"

	"github.com/openkrafter/anytore-backend/logger"
)

var Config *AnytoreConfig

type AnytoreConfig struct {
	GIN_MODE string `yaml:"GIN_MODE"`
	AWS_REGION string `yaml:"AWS_REGION"`
}

func newConfig() (*AnytoreConfig, error) {
	config := new(AnytoreConfig)

	// env: GIN_MODE, value: debug or release, default: debug
	config.GIN_MODE = "debug"
	if os.Getenv("GIN_MODE") == "debug" || os.Getenv("GIN_MODE") == "release" {
		config.GIN_MODE = os.Getenv("GIN_MODE")
	}

	// env: AWS_REGION, value: ap-northeast-1, us-east-1, ..., default: ap-northeast-1
	config.AWS_REGION = "ap-northeast-1"
	if os.Getenv("AWS_REGION") != "" {
		config.AWS_REGION = os.Getenv("AWS_REGION")
	}

	return config, nil
}

func init() {
	logger.Logger.Debug("Init config.")

	var err error
	Config, err = newConfig()
	if err != nil {
		logger.Logger.Error("Failed to read defaultConfig.yaml.", logger.ErrAttr(err))
	}
	logger.Logger.Debug(Config.GIN_MODE)
	logger.Logger.Debug(Config.AWS_REGION)
	fmt.Printf("type %T", Config.AWS_REGION)

	logger.Logger.Debug("config: ", Config)
}
