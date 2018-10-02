package main

import (
	"fmt"
	"os"
)

func main3() {
	var a int16
	var b, c = "B", 'C'
	var d, e = os.Open("arquivo.x")
	var f float32 = 1.6180339887498948482
	var g = 2.45
	var h = complex(g, 2.2*g)
	var nomes = "abcdef"
	var valores = []interface{}{a, b, c, d, e, f} // declarando um array elastico, slice, to tipo interface{} vazia, aceita qualquer valor, tipo List<Object>
	var (
		i int
		v interface{}
	)
	for i, v = range valores {
		var nome = nomes[i]
		// fmt.Println(i)
		// fmt.Println(v)
		fmt.Printf("var %c %-13T = %#[2]v \n", nome, v)
	}
	fmt.Printf("var %c %-13T = %#[2]v \n", 'g', g)
	fmt.Printf("var %c %-13T = %#[2]v \n", 'h', h)
}
