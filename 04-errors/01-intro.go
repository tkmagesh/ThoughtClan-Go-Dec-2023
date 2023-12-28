package main

import (
	"errors"
	"fmt"
)

var ErrAdd error = errors.New("add operation error")
var ErrDivide error = errors.New("divide operation error")

func main() {
	var err error
	err = errors.New("dummy error")
	// err = ErrDivide
	if err == nil {
		fmt.Println("No errors")
	} else if err == ErrAdd {
		fmt.Println("add operation resulted in an error")
	} else if err == ErrDivide {
		fmt.Println("divide operation resulted in an error")
	} else {
		fmt.Println("Unknown error : ", err.Error())
	}

}
