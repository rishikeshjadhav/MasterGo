package main

import (
	"fmt"
	"math"
	"unsafe"
)

type s1 struct {
	n int
	b bool
}

type s2 struct {
	r []rune
}

type s3 struct {
	r [3]rune
}

type emptyStruct struct{}

type metaEmptyStruct struct {
	A struct{}
	B struct{}
}

type sliceOfNothings []struct{}

func main() {

	fmt.Println("Welcome to comparing structs")

	s11 := s1{n: 4, b: true}
	s12 := s1{n: 4, b: true}

	fmt.Println(s11 == s12) // Does this line compile?

	// s21 := s2{r: []rune{'a', 'b', '🎵'}}
	// s22 := s2{r: []rune{'a', 'b', '🎶'}}

	// fmt.Println(s21 == s22) // Does this line compile?

	s31 := s3{r: [3]rune{'a', 'b', '🎵'}}
	s32 := s3{r: [3]rune{'a', 'b', '🎶'}}

	fmt.Println(s31 == s32) // Does this line compile?

	fmt.Println("Size of emptyStruct:", unsafe.Sizeof(emptyStruct{}))
	fmt.Println("Size of metaEmptyStruct:", unsafe.Sizeof(metaEmptyStruct{}))

	sOfN := make(sliceOfNothings, math.MaxInt64)
	fmt.Println("Size of sOfN:", unsafe.Sizeof(sOfN))

	voidSlice := make([]struct{}, 1000, 5000)
	fmt.Println("Size of voidSlice:", unsafe.Sizeof(voidSlice))
}
