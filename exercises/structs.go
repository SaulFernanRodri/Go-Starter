package main

import "fmt"

// Definición de una Struct
type Men struct {
	name string
	age  int
}

// Método de la Struct
func (p Men) greet() {
	fmt.Println("Hello, my name is", p.name)
}

func structs() {
	// Crear una instancia de Person
	person1 := Men{name: "Álvaro", age: 22}
	fmt.Println("Person details:", person1)

	// Llamar al método greet
	person1.greet()
}
