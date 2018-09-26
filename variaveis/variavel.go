package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e é publico
	Endereco string
	//Telefone é um valor importante para o curso
	Telefone string
	// Variavel Endereco é publica e telefone é privada, só pelo falo de Endereco começa com letra maiuscula
	quantidade, estoque int     // quantidde = 0
	comprou             bool    // comprou = false
	valor               float64 // valor = 0.00
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("teste: %v\r\n", teste)
}
