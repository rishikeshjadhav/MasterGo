package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	dept := []string{"P & D", "HR", "Admin"}

	for elem := range dept {
		fmt.Println(elem)
	}

	fmt.Println(dept)

}
