package main

import (
	"log"
	"message-queue/order-service/router"
	"net/http"
)

func main() {
	r := router.SetupRouter()
	log.Println("Starting order service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
