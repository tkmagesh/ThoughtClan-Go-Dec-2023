package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- 10 * i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}
