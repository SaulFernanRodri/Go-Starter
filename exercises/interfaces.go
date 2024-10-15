package main

import "fmt"

// Definición de una interfaz
type Speaker interface {
	Speak() string
}

// Struct que implementa la interfaz Speaker
type Worker struct {
	name string
	age  int
}

// Método que implementa la interfaz
func (p Worker) Speak() string {
	return fmt.Sprintf("Hello, I am %s, and I am %d years old.", p.name, p.age)
}

func interfaces() {
	// Crear una instancia de Person
	worker1 := Worker{name: "Álvaro", age: 22}

	// Uso de la interfaz
	var speaker Speaker = worker1
	fmt.Println(speaker.Speak())
}
