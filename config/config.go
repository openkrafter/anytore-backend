package config

import (
	"log/slog"
	"os"

	"github.com/openkrafter/anytore-backend/logger"
)

var Config *AnytoreConfig

type AnytoreConfig struct {
	SampleValue string `yaml:"SAMPLE_VALUE"`
}

func newConfig() *AnytoreConfig {
	config := new(AnytoreConfig)

	// env: SAMPLE_VALUE, value: sample or sample2, default: sample
	config.SampleValue = "sample"
	if os.Getenv("SAMPLE_VALUE") == "sample" || os.Getenv("SAMPLE_VALUE") == "sample2" {
		config.SampleValue = os.Getenv("SAMPLE_VALUE")
	}

	return config
}

func InitConfig() {
	logger.Logger.Debug("Init config.")
	Config := newConfig()
	logger.Logger.Debug("config: ", slog.Any("value ", Config))
}
