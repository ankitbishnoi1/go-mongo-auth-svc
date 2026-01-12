package handlers

import (
	"encoding/json"
	"net/http"

	"mongo-auth-api/internal/service"
)

type AdminHandler struct {
	service *service.DataService
}

func NewAdminHandler(s *service.DataService) *AdminHandler {
	return &AdminHandler{service: s}
}

func (h *AdminHandler) GetOverview(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetAllUserStats()
	if err != nil {
		http.Error(w, "Failed to fetch stats", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stats)
}
