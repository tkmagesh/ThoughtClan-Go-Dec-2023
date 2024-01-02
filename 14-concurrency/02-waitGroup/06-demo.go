package main

import (
	"fmt"
	"sync"
	"time"
)

// communicate by sharing memory
var result int

func main() {

	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = add(100, 200)
	}()

	wg.Wait()
	fmt.Println(result)
}

// 3rd party api
func add(x, y int) int {
	time.Sleep(3 * time.Second)
	return x + y
}
