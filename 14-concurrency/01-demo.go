package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduled to be executed in future
	f2()
	// block the execution so that the scheduler can look for other goroutines that are scheduled and execute them
	time.Sleep(4 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
