package main

import (
	"fmt"
)

func longest(values ...string) int {

	length := 0

	for index, item := range values {

		fmt.Println(index, item)
		currentLength := len(item)
		fmt.Println("Current Length is:", currentLength)
		fmt.Println("Length is:", length)

		if length < currentLength {
			fmt.Printf("\nUpdating length from %d to %d\n", length, currentLength)
			length = currentLength
		}
	}

	fmt.Println("Final Length is", length)

	return length
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards", "aa"))
}
