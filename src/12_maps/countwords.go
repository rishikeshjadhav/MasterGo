package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string, words map[string]int) {

	var splittedWords []string = strings.Split(sentence, " ")

	fmt.Println("Splitted Words: ", splittedWords)
	fmt.Println("Sentence: ", sentence)
	fmt.Println("Words: ", words)

	for _, element := range splittedWords {

		element = strings.Trim(element, " \t\n\"'.,:;?!()-.")

		element = strings.ToLower(element)

		// fmt.Printf("\n%d -> %s\n", index, element)

		words[element]++

	}

	fmt.Println(words)

}

func main() {
	fmt.Println("Welcome to word counter")

	sentence := "This senTeNce is uSed senTeNce uSed to cOunT senTeNce WorDs  WorDsone more"
	words := map[string]int{}

	fmt.Println("Sentence: ", sentence)
	fmt.Println("Words: ", words)

	countWords(sentence, words)

}
