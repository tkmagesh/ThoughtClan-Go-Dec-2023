/* iota - auto generated constant values */
package main

import "fmt"

func main() {
	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	const (
		red   int = iota
		green int = iota
		blue  int = iota
	)
	fmt.Printf("red = %d, green = %d and blue = %d\n", red, green, blue)
}
