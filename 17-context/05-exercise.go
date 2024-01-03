/*
- Write a program to find the prime numbers within the given start and end
- The prime numbers need to be generated concurrently
- Print the generated prime numbers in the main() function as and when they are generated
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type PrimeResult struct {
	No      int
	IsPrime bool
}

var started, processed, cancelled int64

func main() {
	rootCtx := context.Background()

	timeoutCtx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	ch := genPrimes(timeoutCtx, 3, 100)
	for primeNo := range ch {
		fmt.Println("[main] primeNo :", primeNo)
	}
	fmt.Printf("started : %d, processed : %d\n, cancelled : %d\n", started, processed, cancelled)
}

func genPrimes(ctx context.Context, start, end int) <-chan int {
	ch := make(chan int)
	primeCh := make(chan PrimeResult)
	go func() {
		wg := &sync.WaitGroup{}
		for i := start; i <= end; i++ {
			wg.Add(1)
			started++
			go isPrime(ctx, i, primeCh, wg)
		}
		go func() {
			wg.Wait()
			close(primeCh)
		}()

		for result := range primeCh {
			atomic.AddInt64(&processed, 1)
			fmt.Println("genPrimes :", result)
			if result.IsPrime {
				ch <- result.No
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(ctx context.Context, no int, primeCh chan<- PrimeResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[isPrime] Prime calculation started for no :", no)
	for i := 2; i <= (no / 2); i++ {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			atomic.AddInt64(&cancelled, 1)
			fmt.Println("[isPrime] Cancellation siganl received for no :", no)
			return
		default:
			if no%i == 0 {
				primeCh <- PrimeResult{No: no, IsPrime: false}
				return
			}
		}
	}
	primeCh <- PrimeResult{No: no, IsPrime: true}
}
