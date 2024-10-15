package main

import "fmt"

func maps() {
	// Crear un mapa para almacenar edades de personas
	edades := make(map[string]int)

	// Agregar elementos al mapa
	edades["Juan"] = 25
	edades["Ana"] = 30
	edades["Luis"] = 35

	// Mostrar el mapa completo
	fmt.Println("Mapa completo:", edades)

	// Acceder a un valor específico
	edadJuan := edades["Juan"]
	fmt.Println("Edad de Juan:", edadJuan)

	// Verificar si una clave existe en el mapa
	if edad, existe := edades["Pedro"]; existe {
		fmt.Println("Edad de Pedro:", edad)
	} else {
		fmt.Println("Pedro no está en el mapa")
	}

	// Eliminar un elemento del mapa
	delete(edades, "Ana")
	fmt.Println("Mapa después de eliminar a Ana:", edades)

	// Iterar sobre el mapa
	fmt.Println("Iterando sobre el mapa:")
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}
}
