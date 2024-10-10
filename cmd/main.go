// cmd/mi-aplicacion/main.go
package main

import (
	"github.com/SaulFernanRodri/go-starter/internal/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/usuarios", handlers.UserHandler)
	http.HandleFunc("/productos", handlers.ProductHandler)

	http.ListenAndServe(":8080", nil)
}
