/* Handling multiple channels (using select-case)*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 200
	}()

	ch3 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		d3 := <-ch3
		fmt.Println("Data from ch3 :", d3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
		case d2 := <-ch2:
			fmt.Println(d2)
		default:
			fmt.Println("No channel operations were successful")
		}
	}

}
