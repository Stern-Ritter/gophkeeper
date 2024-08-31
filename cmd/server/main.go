package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	app "github.com/Stern-Ritter/gophkeeper/internal/app/server"
	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	printBuildInfo()

	config, err := app.GetConfig(config.ServerConfig{
		URL:             "localhost:3300",
		ShutdownTimeout: 5,
		LoggerLvl:       "debug",
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}

	logger, err := logger.Initialize(config.LoggerLvl)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = app.Run(&config, logger)
	if err != nil {
		logger.Fatal("Failed to start server", zap.String("event", "start server"),
			zap.Error(err))
	}
}

func printBuildInfo() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)
}
