package main

import (
	"fmt"

	"github.com/rishikeshjadhav/MasterGo/src/09_packages/calculator"
)

func main() {
	fmt.Println("Welcome to using calculator")

	// a, b, c int

	a, b, c := 3, 5, 0

	c = calculator.Add(a, b)

	fmt.Printf("Addition of %d and %d is %d", a, b, c)
}
