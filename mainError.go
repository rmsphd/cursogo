package main

import (
	"fmt"

	"github.com/rmsphd/cursogo/matematica"
)

func mainError() {
	resultado, err := matematica.Calculo(matematica.Multiplicador, 2, 2)
	if HandlerError(err) {
		panic(err)
	}
	fmt.Println("resultado: ", resultado)

}
