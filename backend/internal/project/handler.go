package project

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
	userID := auth.GetUserID(r.Context())

	var req struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	p, err := h.service.Create(r.Context(), req.Name, req.Description, userID)
	if err != nil {
		// http.Error(w, err.Error(), 400)
		response.Error(w, 400, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(p)
	response.JSON(w, http.StatusCreated, p)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r.Context())

	projects, err := h.service.GetAll(r.Context(), userID)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		response.Error(w, 500, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(projects)
	response.JSON(w, http.StatusOK, projects)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	p, err := h.service.GetByID(r.Context(), id, userID)
	if err != nil {
		// http.Error(w, err.Error(), 403)
		response.Error(w, 403, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(p)
	response.JSON(w, http.StatusOK, p)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := auth.GetUserID(r.Context())

	var req struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	p := &Project{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	err := h.service.Update(r.Context(), p, userID)
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
