package main

import (
	"net/http"

	"github.com/SaulFernanRodri/go-starter/internal/handlers"
)

func main() {

	http.HandleFunc("/usuarios", handlers.UserHandler)
	http.HandleFunc("/productos", handlers.ProductHandler)
	http.ListenAndServe(":8080", nil)
}
