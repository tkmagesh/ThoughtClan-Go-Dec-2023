package main

import "fmt"

func main() {
	// invoke f1 through "exec" fn
	exec(f1)

	// invoke f2 through "exec" fn
	exec(f2)

	exec(f3)

	exec(func() {
		fmt.Println("anon fn invoked")
	})

}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3() {
	fmt.Println("f3 invoked")
}
