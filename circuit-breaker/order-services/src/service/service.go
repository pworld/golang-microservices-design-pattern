package service

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fail := rand.Intn(10) < 3 // 30% failure rate

	if fail {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	order := Order{ID: rand.Intn(1000), Status: "Processing"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
