package main

import (
	"fmt"
	"os"
	"strings"
)

func acronym(s string) (acr string) {

	// TODO: Your code here

	for _, item := range strings.Split(s, " ") {
		// fmt.Println(index, "  ", string(item[0]))
		acr += string(item[0])
	}

	return acr
}

func main() {

	a := -1.0
	b := 0.0

	fmt.Println(a / b)

	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
