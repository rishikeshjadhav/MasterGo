package main

import (
	"fmt"
)

type IntSet map[int]struct{}

// iter returns a slice of empty structs of length n.
func iter(n int) []struct{} {
	return make([]struct{}, n)
}

// Now we can run a range loop over the return value of iter().
func iterateFrom1ToN(n int) {
	for i := range iter(n) {
		fmt.Print(i, ",")
	}
	fmt.Println()
}

func main() {

	fmt.Println("Welcome to Set using struct")

	set := IntSet{}

	var el = struct{}{}

	set[5] = el
	set[28] = el
	set[12] = el
	set[12] = el // adding twice does not change the set
	fmt.Println("Set after adding 5, 28, 12, and 12:", set)

	// Test if a value exists in the set

	if _, ok := set[5]; ok {
		fmt.Println("5 exists")
	}

	// Get all values from the set
	fmt.Print("'set' consists of: ")
	for element := range set {
		fmt.Print(element, ",")
	}
	fmt.Println()

	iterateFrom1ToN(7)
}
