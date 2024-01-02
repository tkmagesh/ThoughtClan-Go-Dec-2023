package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := add(100, 200)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
