package calculator

import "fmt"

var opCount map[string]int

func init() {
	opCount = make(map[string]int)
	fmt.Println("calculator initialized - [calc.go]")
}

func OpCount() map[string]int {
	return opCount
}
