package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"go-microservice/internal/app"
	"go-microservice/internal/config"
)

var (
	appname = "go-microservice"
	version = "1.0.0"
	build   = "20240915"
	public  = "0.0.0.0:8080"
	private = "0.0.0.0:8081"
	debug   = false
)

func init() {
	Logger := zap.NewExample()
	defer Logger.Sync()
	Logger.Info("init",
		zap.String("App version:", version),
		zap.String("App build:", build),
		zap.String("App name:", appname),
		zap.String("Public http at:", "http://"+public),
		zap.String("Private http at:", "http://"+private),
	)
}

func main() {
	// metrics.RegisterMetrics()

	cfg := config.NewAppConfig(version)
	logger := logging.NewAppLogger(cfg)

	a := app.NewApp(cfg, logger)

	a.Run()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	a.Stop()
}
