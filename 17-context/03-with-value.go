package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// top most context
	rootCtx := context.Background()

	// context with value
	valCtx := context.WithValue(rootCtx, "root-key", "roov-value")

	// context with cancellation by timeout
	cancelCtx, cancel := context.WithTimeout(valCtx, 10*time.Second)

	fmt.Println("Will stop after 10 seconds.. hit ENTER to stop manually...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printTime(cancelCtx, wg)

	wg.Add(1)
	go printFib(cancelCtx, wg)

	wg.Wait()
	fmt.Println(cancelCtx.Err())
}

func printTime(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[printTime], root-key :", ctx.Value("root-key"))
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
	fmt.Println("[printFib], root-key :", ctx.Value("root-key"))
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
