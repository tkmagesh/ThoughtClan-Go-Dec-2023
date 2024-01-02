/*
- Write a program to find the prime numbers within the given start and end
- The prime numbers need to be generated concurrently
- Print the generated prime numbers in the main() function as and when they are generated
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes(3, 100)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			if isPrime(i) {
				time.Sleep(500 * time.Millisecond)
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
