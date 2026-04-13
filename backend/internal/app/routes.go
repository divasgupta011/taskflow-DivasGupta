package app

import (
	"net/http"
)

func (a *App) routes() {
	a.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	a.router.Post("/auth/register", a.authHandler.Register)
	a.router.Post("/auth/login", a.authHandler.Login)

}
