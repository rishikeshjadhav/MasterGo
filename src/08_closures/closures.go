package main

import "fmt"

func newClosures() (func(), func() int) {
	a := 0
	// Your code here!

	fmt.Println("Inside newClosures with a as ", a)

	f1 := func() {
		fmt.Println("Inside f1 with a as ", a)
		a = 5
		fmt.Println("After f1 with a as ", a)
	}

	f2 := func() int {
		fmt.Println("Inside f2 with a as ", a)
		a = a * 7
		fmt.Println("After f2 with a as ", a)
		return a
	}

	return f1, f2

}

func main() {
	f1, f2 := newClosures()
	f1()      // sets "a" to 5
	n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call?
	fmt.Println(n)
}
