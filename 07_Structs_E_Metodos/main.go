package main

import "fmt"

// Definindo uma struct
type Pessoa struct {
	nome  string
	idade int
}

// Método associado a struct Pessoa
func (p Pessoa) saudacao() {
	fmt.Println("Olá,", p.nome)
}

func main() {
	pessoa := Pessoa{nome: "Maria", idade: 30}
	pessoa.saudacao() // Chamando o método
}
