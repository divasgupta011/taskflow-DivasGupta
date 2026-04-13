package app

import (
	"database/sql"
	"net/http"
	"taskflow/internal/auth"
	"taskflow/internal/config"
	"taskflow/internal/db"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router      *chi.Mux
	cfg         *config.Config
	db          *sql.DB
	authHandler *auth.Handler
}

func New() *App {
	r := chi.NewRouter()

	cfg := config.Load()

	db.RunMigrations(cfg.DBUrl())

	database, err := db.New(cfg)
	if err != nil {
		panic(err)
	}

	authRepo := auth.NewRepository(database)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	app := &App{
		router:      r,
		cfg:         cfg,
		db:          database,
		authHandler: authHandler,
	}

	app.routes()
	return app
}

func (a *App) Router() http.Handler {
	return a.router
}
