package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to flags")

	var max int
	flag.IntVar(&max, "max", 999, "Defines max value")

	verbose := flag.Bool("verbose", false, "Define type of loggin required")
	count := flag.Int("count", 10, "Defines count")

	flag.Parse()

	fmt.Println("Versbose flag is ", *verbose)
	fmt.Println("Count flag is ", *count)
	fmt.Println("Count flag is ", max)

	fmt.Println(os.Args)
}
