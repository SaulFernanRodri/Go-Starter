package main

import (
	"net/http"

	"go-starter/web/internal/handlers"
)

func main() {

	http.HandleFunc("/usuarios", handlers.UserHandler)
	http.HandleFunc("/productos", handlers.ProductHandler)
	http.ListenAndServe(":8080", nil)
}
