package app

import (
	"net/http"

	"taskflow/internal/auth"

	"github.com/go-chi/chi/v5"
)

// func (a *App) routes() {
// 	a.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("ok"))
// 	})

// 	a.router.Post("/auth/register", a.authHandler.Register)
// 	a.router.Post("/auth/login", a.authHandler.Login)

// 	a.router.Group(func(r chi.Router) {
// 		r.Use(auth.AuthMiddleware(a.cfg.JWTSecret))

// 		r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
// 			userID := auth.GetUserID(r.Context())
// 			w.Write([]byte("hello user " + userID))
// 		})

// 		r.Route("/projects", func(r chi.Router) {
// 			r.Get("/", a.projectHandler.GetAll)
// 			r.Post("/", a.projectHandler.Create)
// 			r.Get("/{id}", a.projectHandler.GetByID)
// 			r.Patch("/{id}", a.projectHandler.Update)
// 			r.Delete("/{id}", a.projectHandler.Delete)
// 		})

// 		r.Route("/projects/{id}/tasks", func(r chi.Router) {
// 			r.Get("/", a.taskHandler.GetByProject)
// 			r.Post("/", a.taskHandler.Create)
// 		})

// 		r.Route("/tasks", func(r chi.Router) {
// 			r.Patch("/{id}", a.taskHandler.Update)
// 			r.Delete("/{id}", a.taskHandler.Delete)
// 		})

// 	})

// }

func (a *App) routes() {
	a.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	a.router.Post("/auth/register", a.authHandler.Register)
	a.router.Post("/auth/login", a.authHandler.Login)

	a.router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(a.cfg.JWTSecret))

		r.Route("/projects", func(r chi.Router) {
			r.Get("/", a.projectHandler.GetAll)
			r.Post("/", a.projectHandler.Create)
			r.Get("/{id}", a.projectHandler.GetByID)
			r.Patch("/{id}", a.projectHandler.Update)
			r.Delete("/{id}", a.projectHandler.Delete)

			r.Route("/{id}/tasks", func(r chi.Router) {
				r.Get("/", a.taskHandler.GetByProject)
				r.Post("/", a.taskHandler.Create)
			})
		})

		r.Route("/tasks", func(r chi.Router) {
			r.Patch("/{id}", a.taskHandler.Update)
			r.Delete("/{id}", a.taskHandler.Delete)
		})
	})

}
