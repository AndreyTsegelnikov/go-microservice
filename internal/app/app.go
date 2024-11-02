package app

import (
	"go-microservice/internal/config"
	"time"
	"go-microservice/internal/logger"
)

type App struct {
	cfg    *config.AppConfig
	logger *logger.AppLogger
	srv    *server.Server
}

func New(cfg *config.AppConfig, logger *logging.AppLogger) *App {
	promoProvider, err := httpcli.NewFastClient(cfg.MarketingHosts, cfg, logger, httpcli.DefaultFConfig("promo"), httpcli.DefaultCBConfig())
	if err != nil {
		logger.Fatal(err)
	}

	promoRepositoryV1 := repository.NewPromoRepository[*v1.PromoResponse](cfg, logger, promoProvider)
	promoRepositoryV2 := repository.NewPromoRepository[*v2.PromoResponse](cfg, logger, promoProvider)
	promoService := service.NewPromoService(promoRepositoryV1, promoRepositoryV2)

	srv := server.NewServer(cfg, logger, promoRepositoryV1, promoRepositoryV2, promoService)

	return &App{
		cfg:    cfg,
		logger: logger,
		srv:    srv,
	}
}

func (app *App) Run() {
	app.logger.Info("Starting app: " + app.cfg.AppName)
	app.logger.Info("App version: " + app.cfg.AppVersion)
	app.logger.Info("App identity: " + app.cfg.AppId)

	app.srv.Run()

	app.logger.Info("App started: " + time.Now().String())

	metrics.ReplicCountByVersion.WithLabelValues(app.cfg.AppVersion).Inc()
}

func (app *App) Stop() {
	app.logger.Info("Stopping " + app.cfg.AppName + "...")

	app.srv.Stop()

	app.logger.Info("App stopped: " + time.Now().String())
}
