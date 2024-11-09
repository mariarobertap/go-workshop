package main

import (
	"fmt"
	"math" // Pacote da biblioteca padrão
)

func main() {
	raio := 5.0
	area := math.Pi * math.Pow(raio, 2)

	fmt.Println("Área do círculo com raio 5:", area)
}
