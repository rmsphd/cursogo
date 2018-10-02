package main

import (
	"fmt"
	"unicode/utf8"
)

func main_s() {
	ss := []string{"bola", "café", "世界", "🌎"}
	for _, s := range ss {
		fmt.Println("%s %d %d %t", s, len(s),
			utf8.RuneCountInString(s), []rune(s))
	}
}
