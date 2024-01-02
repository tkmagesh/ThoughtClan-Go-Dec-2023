/* simulating deadlock */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(10) // increment the wg counter by 1
	go f1(wg)  //scheduled to be executed in future
	f2()
	// block the execution so that the scheduler can look for other goroutines that are scheduled and execute them
	wg.Wait() // block until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
