package main

import (
	"log"
	"net/http"

	"go-starter/web/internal/handlers"
)

func main() {

	http.HandleFunc("/usuarios", handlers.UserHandler)
	http.HandleFunc("/productos", handlers.ProductHandler)
	http.HandleFunc("/showimage", handlers.ShowImage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
