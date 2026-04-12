package app

import (
	"net/http"
	"taskflow/internal/config"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router *chi.Mux
	cfg    *config.Config
}

func New() *App {
	r := chi.NewRouter()

	cfg := config.Load()

	app := &App{
		router: r,
		cfg:    cfg,
	}

	app.routes()
	return app
}

func (a *App) Router() http.Handler {
	return a.router
}
