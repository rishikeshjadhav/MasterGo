package token

import (
	"fmt"
	"strings"
)

func Scan(s string) []string {
	// return strings.Split(s, " ")
	return strings.Fields(s)
}

func main() {

	fmt.Println("Welcome to unit testing")

	s := "This is a tokenised string for testing "

	tokens := Scan(s)

	fmt.Println(s)
	fmt.Println(tokens)
	fmt.Println(len(tokens))
}
