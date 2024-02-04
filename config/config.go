package config

import (
	"fmt"
	"os"

	"github.com/openkrafter/anytore-backend/logger"
	"gopkg.in/yaml.v3"
)

var Config *AnytoreConfig

type AnytoreConfig struct {
	GIN_MODE string `yaml:"GIN_MODE"`
	AWS_REGION string `yaml:"AWS_REGION"`
}

func newConfig() (*AnytoreConfig, error) {
	config := new(AnytoreConfig)

	defaultConfig, err := os.ReadFile("config/defaultConfig.yaml")
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(defaultConfig, &config)

	if os.Getenv("GIN_MODE") != "" {
		config.GIN_MODE = os.Getenv("GIN_MODE")
	}

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
