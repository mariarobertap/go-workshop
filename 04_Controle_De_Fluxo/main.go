package main

import "fmt"

func main() {
	numero := 10

	// Condicional if
	if numero > 5 {
		fmt.Println("Número maior que 5")
	} else {
		fmt.Println("Número menor ou igual a 5")
	}

	// Laço for
	for i := 0; i < 5; i++ {
		fmt.Println("Iteração:", i)
	}

	// Switch case
	switch numero {
	case 10:
		fmt.Println("O número é 10")
	default:
		fmt.Println("Outro número")
	}
}
