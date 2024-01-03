package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printTime(cancelCtx, wg)

	wg.Add(1)
	go printFib(cancelCtx, wg)

	fmt.Println("hit ENTER to stop...")
	fmt.Scanln()
	cancel()
	wg.Wait()

}

func printTime(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[printTime] Cancel signal received.. stopping")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now())
		}
	}
}

func printFib(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for x, y := 0, 1; ; {
		select {
		case <-ctx.Done():
			fmt.Println("[printFib] Cancel signal received.. stopping")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Fib No :", x)
			x, y = y, x+y
		}
	}
}
