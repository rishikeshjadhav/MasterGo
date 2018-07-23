package main

import "fmt"

func main() {
	// Print RGB values...
	r, g, b := 124, 87, 3
	// ...as #7c5703  (specifying hex format, fixed width, and leading zeroes)
	// Hint: don't forget to add a newline at the end of the format string.
	fmt.Printf("#%02x%02x%02x", r, g, b)
	// ...as rgb(124, 87, 3)
	fmt.Printf("\nrgb(%d, %d, %d)", r, g, b)
	// ...as rgb(124, 087, 003) (specifying fixed width and leading zeroes)
	fmt.Printf("\nrgb(%03d, %03d, %03d)", r, g, b)
	// ...as rgb(48%, 34%, 1%) (specifying a literal percent sign)
	fmt.Printf("\nrgb(%d%%, %d%%, %d%%)", r, g, b)
	// Print the type of r.
	fmt.Printf("\n%T", r)
}
