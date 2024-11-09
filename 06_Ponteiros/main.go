package main

import "fmt"

func main() {
	a := 58
	b := &a // Ponteiro para a variável a

	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor do ponteiro b:", *b)
}
