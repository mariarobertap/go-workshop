package main

import "fmt"

func main() {
	// Array de 3 elementos
	var numeros [3]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	// Slice
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Array:", numeros)
	fmt.Println("Slice:", slice)
}
