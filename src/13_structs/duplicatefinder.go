package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Welcome to duplicate finder")

	if len(os.Args) < 2 {
		fmt.Println("Please provide path to file")
		return
	}

	fmt.Printf("Reading from file %s\n", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	set := make(map[string]struct{})

	s := bufio.NewScanner(f)

	for s.Scan() {
		// The Text() method returns the current token as a string.
		if _, ok := set[s.Text()]; ok {
			fmt.Println("Duplicate line:", s.Text())
		} else {
			set[s.Text()] = struct{}{}
		}
	}

	fmt.Printf("\nSet: %v", set)

	fmt.Println("\nCompleted duplicate finder")
}
