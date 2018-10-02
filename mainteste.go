package main

import (
	"fmt"

	"github.com/rmsphd/cursogo/matematica"
)

func mainTeste() {
	resultado, _ := matematica.Calculo(matematica.Multiplicador, 2, 2)
	// resultado = matematica.Soma(resultado, 2)
	fmt.Printf("Resultado da multiplicacao %d \r\n", resultado)

	fmt.Println("Ola mundo go")
	fmt.Println("ola $s", "2")
}
