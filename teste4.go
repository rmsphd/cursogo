package main

import (
	"fmt"
)

func mainteste4() {

	m0 := map[string]int{}

	m0["k1"] = 7
	m0["k2"] = 13

	fmt.Println("m0", m0)
	fmt.Println("len", len(m0))

	v2, prs1 := m0["k2"]
	fmt.Println("k2 v2", v2, prs1)

	delete(m0, "k2")

	fmt.Println(m0)
	fmt.Println("len", len(m0))

	v3 := m0["k2"]
	fmt.Println("k2 v3", v3)

	v4, prs := m0["k2"]
	fmt.Println("k2 v4", v4, prs)

	m0["k2"] = 0

	v5, prs2 := m0["k2"]
	fmt.Println("k2 v5", v5, prs2)

	v6, prs3 := m0["k3"]
	fmt.Println("k3", v6, prs3)

}
