package app

import (
	"database/sql"
	"net/http"
	"taskflow/internal/config"
	"taskflow/internal/db"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router *chi.Mux
	cfg    *config.Config
	db     *sql.DB
}

func New() *App {
	r := chi.NewRouter()

	cfg := config.Load()

	database, err := db.New(cfg)
	if err != nil {
		panic(err)
	}

	app := &App{
		router: r,
		cfg:    cfg,
		db:     database,
	}

	app.routes()
	return app
}

func (a *App) Router() http.Handler {
	return a.router
}
