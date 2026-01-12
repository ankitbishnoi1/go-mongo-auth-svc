package router

import (
	"mongo-auth-api/internal/config"
	"mongo-auth-api/internal/handlers"
	"mongo-auth-api/internal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(
	authHandler *handlers.AuthHandler,
	dataHandler *handlers.DataHandler,
	adminHandler *handlers.AdminHandler,
	cfg *config.Config,
) *mux.Router {
	r := mux.NewRouter()

	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	protected := api.PathPrefix("/data").Subrouter()
	protected.Use(authMiddleware.Handle)
	protected.HandleFunc("", dataHandler.GetData).Methods("GET")

	admin := api.PathPrefix("/admin").Subrouter()
	// In a real app, admin routes would have separate/stricter middleware
	admin.Use(authMiddleware.Handle)
	admin.HandleFunc("/overview", adminHandler.GetOverview).Methods("GET")

	return r
}
