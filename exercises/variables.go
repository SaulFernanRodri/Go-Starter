package main

import "fmt"

func variables() {
	// Declaración de variables
	var entero = 10
	var flotante = 20.5
	var cadena = "Hola, Go!"
	var booleano = true

	// Declaración de variables sin inicialización
	var sinInicializar int

	// Declaración corta de variables
	enteroCorto := 15
	flotanteCorto := 25.5
	cadenaCorto := "Hola, Mundo!"
	booleanoCorto := false

	// Declaración de arrays
	var arrayEnteros = [5]int{1, 2, 3, 4, 5}
	var arrayCadenas = [3]string{"uno", "dos", "tres"}

	// Imprimir variables
	/*
		%s para cadenas.
		%d para enteros.
		%f para números de punto flotante.
		%T para imprimir el tipo del valor.
		%t para valores booleanos.
		%v para imprimir el valor en formato predeterminado.
	*/

	fmt.Printf("Entero: %d\n", entero)
	fmt.Printf("Flotante: %.2f\n", flotante)
	fmt.Printf("Cadena: %s\n", cadena)
	fmt.Printf("Booleano: %t\n", booleano)
	fmt.Printf("Sin inicializar: %d\n", sinInicializar)
	fmt.Printf("Entero corto: %d\n", enteroCorto)
	fmt.Printf("Flotante corto: %.2f\n", flotanteCorto)
	fmt.Printf("Cadena corta: %s\n", cadenaCorto)
	fmt.Printf("Booleano corto: %t\n", booleanoCorto)

	// Imprimir arrays
	fmt.Printf("Array de enteros: %v\n", arrayEnteros)
	fmt.Printf("Array de cadenas: %v\n", arrayCadenas)

	// Lectura de datos desde la entrada estándar
	var entrada string
	fmt.Println("Introduce un valor: ")
	if _, err := fmt.Scanln(&entrada); err != nil {
		fmt.Println("Error al leer la entrada:", err)
	}
	fmt.Printf("Valor introducido: %s\n", entrada)

	// Operaciones aritméticas
	suma := entero + enteroCorto
	resta := flotante - flotanteCorto
	multiplicacion := entero * enteroCorto
	division := flotante / flotanteCorto

	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Resta: %.2f\n", resta)
	fmt.Printf("Multiplicación: %d\n", multiplicacion)
	fmt.Printf("División: %.2f\n", division)

	// Operaciones lógicas
	and := booleano && booleanoCorto
	or := booleano || booleanoCorto
	not := !booleano

	fmt.Printf("AND: %t\n", and)
	fmt.Printf("OR: %t\n", or)
	fmt.Printf("NOT: %t\n", not)
}
