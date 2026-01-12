package main

import (
	"net/http"

	"mongo-auth-api/internal/config"
	"mongo-auth-api/internal/database"
	"mongo-auth-api/internal/handlers"
	"mongo-auth-api/internal/repository"
	"mongo-auth-api/internal/router"
	"mongo-auth-api/internal/service"
	"mongo-auth-api/pkg/logger"
)

func main() {
	logger.Init()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}

	if err := database.Connect(cfg.MongoURI); err != nil {
		logger.Fatal("Failed to connect to MongoDB: ", err)
	}

	userRepo := repository.NewUserRepository()

	authService := service.NewAuthService(userRepo, cfg)
	dataService := service.NewDataService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	dataHandler := handlers.NewDataHandler(dataService)
	adminHandler := handlers.NewAdminHandler(dataService)

	r := router.NewRouter(authHandler, dataHandler, adminHandler, cfg)

	logger.Info("Starting server on ", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		logger.Fatal("Server failed: ", err)
	}
}
