package main

import (
	"log"

	"github.com/rivo/tview"
	"go.uber.org/zap"

	app "github.com/Stern-Ritter/gophkeeper/internal/app/client"
	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
	service "github.com/Stern-Ritter/gophkeeper/internal/service/client"
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

	l, err := logger.Initialize(cfg.LoggerLvl)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	c := service.NewClient(&cfg)
	a := tview.NewApplication()
	err = app.Run(c, a, &cfg, l)
	if err != nil {
		l.Fatal(err.Error(), zap.String("event", "start application"))
	}
}
