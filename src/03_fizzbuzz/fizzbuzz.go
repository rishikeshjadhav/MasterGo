package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	// Your code here!

	fmt.Println("Received n as ", n)

	for i := 1; i <= n; i++ {

		switch {
		case i%15 == 0:
			fmt.Println(i, ": FizzBuzz")
			fallthrough
		case i%5 == 0:
			fmt.Println(i, ": Buzz")
			fallthrough
		case i%3 == 0:
			fmt.Println(i, ": Fizz")
		default:
			fmt.Println(i, ": Not divisible")
		}

	}

}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
