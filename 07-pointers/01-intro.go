package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no // value -> address (pointer)
	fmt.Println(noPtr, no)

	// dereferencing
	// address (pointer) -> value

	x := *noPtr
	fmt.Println("x :", x)

	*noPtr += 100
	fmt.Println(no)

	fmt.Println(no == *(&no))
}
