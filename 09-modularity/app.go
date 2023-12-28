package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/thoughclan-go-dec-2023/09-modularity/calculator"
	// "github.com/tkmagesh/thoughclan-go-dec-2023/09-modularity/calculator/utils"

	ut "github.com/tkmagesh/thoughclan-go-dec-2023/09-modularity/calculator/utils"
)

func run() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	color.Red("modularity app executed")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.Multiply(100, 200))
	fmt.Println(calculator.Divide(200, 0))
	fmt.Println("Operation Count :", calculator.OpCount())

	// fmt.Println("IsEven(21) :", utils.IsEven(21))
	fmt.Println("IsEven(21) :", ut.IsEven(21))
	return
}
