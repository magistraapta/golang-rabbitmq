package handler

import (
	"encoding/json"
	"message-queue/order-service/models"
	"message-queue/order-service/publisher"
	"net/http"

	"github.com/google/uuid"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()

	if err := publisher.PublisherOrder(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order created successfully"})
}
