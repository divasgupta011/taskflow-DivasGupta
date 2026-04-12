package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router *chi.Mux
}

func New() *App {
	r := chi.NewRouter()

	app := &App{
		router: r,
	}

	app.routes()
	return app
}

func (a *App) Router() http.Handler {
	return a.router
}
