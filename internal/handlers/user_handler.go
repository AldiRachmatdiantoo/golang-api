package handlers

import (
	"net/http"
)

type userHandler interface {
	GetUsersData()
}
type Handler struct {
	svc userHandler
}

func NewHandler(services userHandler) *Handler {
	return &Handler{svc: services}
}
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	am := "Aman aja"
	w.Write([]byte(am))
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) Store(w http.ResponseWriter, r *http.Request)   {}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) Destroy(w http.ResponseWriter, r *http.Request) {}
