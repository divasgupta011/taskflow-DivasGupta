package task

import (
	"encoding/json"
	"net/http"

	"taskflow/internal/auth"
	"taskflow/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	var t Task
	json.NewDecoder(r.Body).Decode(&t)
	t.ProjectID = projectID

	err := h.service.Create(r.Context(), &t, userID)
	if err != nil {
		// http.Error(w, err.Error(), 403)
		response.Error(w, 403, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(t)
	response.JSON(w, http.StatusCreated, t)
}

func (h *Handler) GetByProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	status := r.URL.Query().Get("status")
	assignee := r.URL.Query().Get("assignee")

	tasks, err := h.service.GetByProject(r.Context(), projectID, status, assignee, userID)
	if err != nil {
		// http.Error(w, err.Error(), 403)
		response.Error(w, 403, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(tasks)
	response.JSON(w, http.StatusOK, tasks)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	var t Task
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id

	err := h.service.Update(r.Context(), &t, userID)
	if err != nil {
		// http.Error(w, err.Error(), 403)
		response.Error(w, 403, err.Error())
		return
	}

	// w.WriteHeader(http.StatusOK)
	response.JSON(w, http.StatusOK, map[string]string{"message": "updated"})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	err := h.service.Delete(r.Context(), id, userID)
	if err != nil {
		// http.Error(w, err.Error(), 403)
		response.Error(w, 403, err.Error())
		return
	}

	// w.WriteHeader(http.StatusNoContent)
	response.JSON(w, http.StatusNoContent, nil)
}
