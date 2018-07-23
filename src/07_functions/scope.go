package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "abcde"
	for _, s1 := range s {
		s2 := unicode.ToUpper(s1)
		fmt.Print(string(s2))
	}
	fmt.Println("\n" + s)
}
