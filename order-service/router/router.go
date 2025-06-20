package router

import (
	"message-queue/order-service/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrder).Methods("POST")

	return r
}
