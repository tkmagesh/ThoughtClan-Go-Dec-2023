/*
write a program that will keep generating prime numbers from 3 in 500 millisecond intervals until the user hits the ENTER key
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	primeCh := genPrimes(stopCh)
	for no := range primeCh {
		fmt.Println(no)
	}
	fmt.Println("Done!")
}

func genPrimes(stopCh <-chan struct{}) <-chan int {
	primeCh := make(chan int)
	go func() {
	LOOP:
		for no := 3; ; no++ {
			select {
			case <-stopCh:
				break LOOP
			default:
				if isPrime(no) {
					time.Sleep(500 * time.Millisecond)
					primeCh <- no
				}
			}
		}
		close(primeCh)
	}()
	return primeCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
