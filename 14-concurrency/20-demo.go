package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex
	count int
}

func (sc *SafeCounter) Increment() {
	sc.Lock()
	{
		sc.count++
	}
	sc.Unlock()
}

var sCounter SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(sCounter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sCounter.Increment()
}
