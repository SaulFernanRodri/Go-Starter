package main

import "fmt"

// Función que envía datos a través de un canal
func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Cerramos el canal cuando terminamos de enviar datos
}

func channels() {
	ch := make(chan int)

	// Ejecutamos la goroutine para enviar datos al canal
	go sendData(ch)

	// Recibimos los datos desde el canal
	for val := range ch {
		fmt.Println("Received from channel:", val)
	}
}
