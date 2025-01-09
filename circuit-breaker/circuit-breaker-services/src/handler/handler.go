package handler

import (
	"circuit-breaker-services/src/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetOrders()
	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	response := map[string]string{"message": data}
	json.NewEncoder(w).Encode(response)
}
