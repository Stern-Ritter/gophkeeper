package main

import (
	"log"

	"go.uber.org/zap"

	app "github.com/Stern-Ritter/gophkeeper/internal/app/client"
	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
)

func main() {
	cfg := config.ClientConfig{
		ServerURL:    "localhost:3300",
		BuildVersion: buildVersion,
		BuildDate:    buildDate,
	}

	logger, err := logger.Initialize(cfg.LoggerLvl)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = app.Run(&cfg, logger)
	if err != nil {
		logger.Fatal(err.Error(), zap.String("event", "start application"))
	}
}
