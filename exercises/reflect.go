package main

import (
	"fmt"
	"reflect"
)

// Definimos una struct Producto con diferentes tipos de datos
type Producto struct {
	Nombre    string
	Categoria string
	Precio    float64
	EnStock   bool
	Cantidad  int
}

// Función que usa reflect para inspeccionar los campos y valores de cualquier struct
func inspeccionarCampos(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// Verificamos que el tipo sea una struct
	if typ.Kind() != reflect.Struct {
		fmt.Println("El tipo proporcionado no es una struct")
		return
	}

	// Iteramos sobre cada campo de la struct
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)      // Información sobre el campo
		fieldValue := val.Field(i) // Valor del campo

		// Obtenemos el nombre del campo
		nombre := field.Name

		// Obtenemos el tipo del campo
		tipo := field.Type

		// Obtenemos el valor del campo
		valor := fieldValue.Interface() // .Interface() convierte el valor a un tipo de interfaz genérica

		// Mostramos la información del campo
		fmt.Printf("Campo: %s, Tipo: %s, Valor: %v\n", nombre, tipo, valor)
	}
}

func reflects() {
	// Creamos una instancia de Producto
	producto := Producto{
		Nombre:    "Laptop",
		Categoria: "Electrónica",
		Precio:    1499.99,
		EnStock:   true,
		Cantidad:  50,
	}

	// Llamamos a la función de inspección
	fmt.Println("Inspección de la struct Producto:")
	inspeccionarCampos(producto)
}
