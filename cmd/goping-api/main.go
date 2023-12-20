package main

import (
	"flag"
	"os"

	"github.com/archey347/goping/internal/api"
	"github.com/archey347/goping/internal/config"
	"go.uber.org/zap"
)

func main() {

	logger := zap.Must(zap.NewProduction())

	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}

	defer logger.Sync()

	logger.Info("goping-api initialising")

	var configPath string
	flag.StringVar(&configPath, "config", "/etc/goping/goping-api.yaml", "Path to config file")
	flag.Parse()

	logger.Info("Using config file: " + configPath)

	c, err := config.Load(configPath)

	if err != nil {
		logger.Fatal("Failed to load config: " + err.Error())
	}

	api.Start(c, logger)
}
