package main

import "fmt"

// Função que recebe dois números e retorna a soma
func somar(a int, b int) int {
	return a + b
}

func main() {
	resultado := somar(3, 4)
	fmt.Println("Soma de 3 e 4:", resultado)
}
