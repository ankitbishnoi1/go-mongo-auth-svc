package handlers

import (
	"encoding/json"
	"net/http"

	"mongo-auth-api/internal/service"
)

type DataHandler struct {
	service *service.DataService
}

func NewDataHandler(s *service.DataService) *DataHandler {
	return &DataHandler{service: s}
}

func (h *DataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	data, err := h.service.GetData(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"data": data})
}
