/*
- Write a program to find the prime numbers within the given start and end
- The prime numbers need to be generated concurrently
- Print the generated prime numbers in the main() function as and when they are generated
*/

package main

import (
	"fmt"
	"sync"
)

type PrimeResult struct {
	No      int
	IsPrime bool
}

func main() {
	ch := genPrimes(3, 100)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	primeCh := make(chan PrimeResult)
	go func() {
		wg := &sync.WaitGroup{}
		for i := start; i <= end; i++ {
			wg.Add(1)
			go isPrime(i, primeCh, wg)
		}
		go func() {
			wg.Wait()
			close(primeCh)
		}()

		for result := range primeCh {
			if result.IsPrime {
				ch <- result.No
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int, primeCh chan<- PrimeResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			primeCh <- PrimeResult{No: no, IsPrime: false}
			return
		}
	}
	primeCh <- PrimeResult{No: no, IsPrime: true}
}
