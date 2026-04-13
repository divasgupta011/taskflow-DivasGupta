package app

import (
	"net/http"

	"taskflow/internal/auth"

	"github.com/go-chi/chi/v5"
)

func (a *App) routes() {
	a.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	a.router.Post("/auth/register", a.authHandler.Register)
	a.router.Post("/auth/login", a.authHandler.Login)

	a.router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(a.cfg.JWTSecret))

		r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			userID := auth.GetUserID(r.Context())
			w.Write([]byte("hello user " + userID))
		})
	})

}
