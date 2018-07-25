package main

import "fmt"

func main() {
	src := []int{}
	fmt.Printf("\nsrc: %v", src)
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	fmt.Printf("\nsrc: %v", src)
	dest1 := append(src, 3)
	fmt.Printf("\nsrc: %v", src)
	fmt.Printf("\ndest1: %v", dest1)
	dest2 := append(src, 4)
	// fmt.Println(src, dest1, dest2)
	fmt.Printf("\nsrc: %v", src)
	fmt.Printf("\ndest1: %v", dest1)
	fmt.Printf("\ndest2: %v", dest2)
	fmt.Printf("\nAddress: \nsrc (%p) \ndest1 (%p) \ndest2 (%p)", &src, &dest1, &dest2)
}
