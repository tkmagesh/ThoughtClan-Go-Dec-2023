package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")
var ErrDivideClient error = errors.New("divide client error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			if errors.Is(e, ErrDivideByZero) {
				fmt.Println("[main - recovery] divide by zero error occured")
			}
			if errors.Is(e, ErrDivideClient) {
				fmt.Println("[main - recovery] divide client error occured")
			}
		}
		fmt.Println("Thank you!")
	}()
	for {
		var divisor int
		multiplier := 100
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		q, r, err := divideClient(multiplier, divisor)
		if err != nil {
			if errors.Is(err, ErrDivideByZero) {
				fmt.Println("Please enter a non zero divisor")
				continue
			}
			fmt.Println("Unknown error :", err.Error())
			continue
		}
		fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
		break
	}

}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		fmt.Println("deferred - [divideClient]")
		if e := recover(); e != nil {
			err = e.(error)
			return
		}

	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] - calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
