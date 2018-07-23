package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)

	fmt.Printf("Received %d number of command line arguments\n", len(os.Args))

	for _, param := range os.Args {
		fmt.Println(param)
	}

}
