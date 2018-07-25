package main

import "fmt"

func changeSlice1(s []int) {
	s[0] = 7
}

func changeSlice2(s []int) {
	s = []int{7}
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice1:", s1)
	fmt.Printf("\ns1 header before: %v len (%d) cap (5d)\n", s1, len(s1), cap(s1))
	// changeSlice1(s1)
	changeSlice2(s1)
	fmt.Println("s1 after changeSlice1:", s1)
	fmt.Printf("\ns1 header after: %v len (%d) cap (5d)\n", s1, len(s1), cap(s1))
}
