package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")
var ErrDivideClient error = errors.New("divide client error")

func main() {
	defer func() {
		if err := recover().(error); err != nil {
			if errors.Is(err, ErrDivideByZero) {
				fmt.Println("[main - recovery] divide by zero error occured")
			}
			if errors.Is(err, ErrDivideClient) {
				fmt.Println("[main - recovery] divide client error occured")
			}
		}
		fmt.Println("Thank you!")
	}()
	multiplier := 100
	divisor := 0
	q, r := divideClient(multiplier, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)

}

func divideClient(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("deferred - [divideClient]")
		if err := recover(); err != nil {
			e := fmt.Errorf("%w %w", err.(error), ErrDivideClient)
			panic(e)
		}

	}()
	quotient, remainder = divide(x, y)
	return
}

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
