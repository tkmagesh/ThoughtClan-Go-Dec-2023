package calculator

import "fmt"

func init() {
	fmt.Println("calculator initialized - [subtract.go]")
}

func Subtract(x, y int) int {
	// opCount++
	opCount["subtract"]++
	return x - y
}
