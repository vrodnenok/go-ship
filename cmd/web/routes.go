package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vrodnenok/go-ship/pkg/config"
	"github.com/vrodnenok/go-ship/pkg/handlers"
)

// routes setting up routes
func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
