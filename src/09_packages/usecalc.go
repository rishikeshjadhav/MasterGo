package main

import (
	"calculator"
	"fmt"
)

func main() {
	fmt.Println("Welcome to using calculator")

	a, b, c int

	a = 3
	b = 4

	c = calculator.Add(a, b)

	fmt.Printf("Addition of %d and %d is %d", a, b, c)
}
