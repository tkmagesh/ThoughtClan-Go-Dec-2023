package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum())

	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(sum(nos...))
}

func sum(list ...int) int {
	var result int
	for _, val := range list {
		result += val
	}
	return result
}
