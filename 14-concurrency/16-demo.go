/* Using channels for signaling */

package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutCh := timeout(5 * time.Second)
	ch := genFib(timeoutCh)
	for fibNo := range ch {
		fmt.Println(fibNo)
	}
}

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}

func genFib(timeoutCh chan time.Time) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; {
			select {
			case <-timeoutCh:
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
