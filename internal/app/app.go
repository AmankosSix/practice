package app

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"practice/internal/config"
	delivery "practice/internal/delivery/http"
	"practice/internal/repository"
	"practice/internal/server"
	"practice/internal/service"
	"practice/pkg/database/postgres"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	db, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		logrus.Fatalf("Error when initialize db: %s", err.Error())
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos)
	handlers := delivery.NewHandler(services)

	srv := server.NewServer(cfg, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Info("Server started")
}
