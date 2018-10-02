package main

import (
	"fmt"
)

func Example() {
	main()
	// Output:
	// Olá mundo!
}

func Example_percorrerMap() {
	m := map[int]string{1: "um", 2: "dois", 3: "três"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// Unordered Output:
	// 1 um
	// 2 dois
	// 3 três
}
