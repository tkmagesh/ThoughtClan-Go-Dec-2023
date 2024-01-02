package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)

	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)

	result := <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
