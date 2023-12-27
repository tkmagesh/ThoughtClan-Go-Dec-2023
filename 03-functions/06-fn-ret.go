package main

import "fmt"

func main() {
	var fn func()
	fn = getFn()
	fn()
}

func getFn() func() {
	return func() {
		fmt.Println("anon fn invoked")
	}
}
