package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) int {
	count := 0
	// Your code here!

	fmt.Println("Received n as ", n)

	for n != 1 {
		if n%2 == 0 {
			fmt.Println("N is even for iteration ", count)
			n = n / 2
		} else {
			fmt.Println("N is odd for iteration ", count)
			n = n*3 + 1
		}

		fmt.Println("N is ", n)

		count++
	}

	fmt.Println("Calculated count as ", count)

	return count
}
func main() {
	var n int
	var err error
	if len(os.Args) > 1 { // Read the number from the command line
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // Read the number interactively
		fmt.Println("Input a number > 1: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if n <= 1 {
		fmt.Println("n must be larger than 1.")
		return
	}
	c := collatz(n)
	fmt.Println(n, "requires", c, "steps to reach 1.")
}
