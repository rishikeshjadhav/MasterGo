package main

import (
	"fmt"
	"sort"
)

type List []string

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return len(l[i]) > len(l[j])
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	fmt.Println("Welcome to custom sorter")
	list := List{"really really long", "short", "quite long", "longer"}
	fmt.Printf("\n%#v\n", list)
	sort.Sort(list)
	fmt.Printf("\n%#v\n", list)
}
