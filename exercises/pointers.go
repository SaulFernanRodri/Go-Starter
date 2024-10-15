package main

import "fmt"

// Definición de una Struct
type Person struct {
	name string
	age  int
}

// Método que usa un puntero para modificar la edad
func (p *Person) birthday() {
	p.age += 1
}

func pointers() {
	// Crear una instancia de Person
	person1 := Person{name: "Álvaro", age: 22}

	// Usamos un puntero para modificar el valor original
	p := &person1
	p.birthday() // Incrementa la edad
	fmt.Println("After birthday:", person1.age)
}
