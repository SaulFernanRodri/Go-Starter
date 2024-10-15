package main

import "fmt"

func slices() {

	// Crear un slice vacío
	var s []int
	fmt.Println("Slice vacío:", s)

	// Crear un slice con valores iniciales
	s = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice inicial:", s)

	// Agregar elementos a un slice
	s = append(s, 6, 7)
	fmt.Println("Slice después de append:", s)

	// Crear un slice a partir de un array
	arr := [5]int{10, 20, 30, 40, 50}
	s = arr[1:4]
	fmt.Println("Slice a partir de un array:", s)

	// Copiar un slice a otro
	s2 := make([]int, len(s))
	copy(s2, s)
	fmt.Println("Slice copiado:", s2)

	// Rebanar un slice
	s3 := s[1:3]
	fmt.Println("Rebanar un slice:", s3)

	// Obtener la longitud y capacidad de un slice
	fmt.Println("Longitud del slice:", len(s))
	fmt.Println("Capacidad del slice:", cap(s))

	// Modificar elementos de un slice
	s[0] = 100
	fmt.Println("Slice modificado:", s)

	// Buscar un elemento en un slice
	indice := buscarElemento(s, 30)
	if indice != -1 {
		fmt.Println("Elemento encontrado en el índice:", indice)
	} else {
		fmt.Println("Elemento no encontrado")
	}

	// Eliminar un elemento de un slice
	s = eliminarElemento(s, 2)
	fmt.Println("Slice después de eliminar el elemento en el índice 2:", s)
}

func buscarElemento(slice []int, elemento int) int {
	for i, v := range slice {
		if v == elemento {
			return i
		}
	}
	return -1
}

func eliminarElemento(slice []int, indice int) []int {
	if indice < 0 || indice >= len(slice) {
		return slice
	}
	return append(slice[:indice], slice[indice+1:]...)
}
