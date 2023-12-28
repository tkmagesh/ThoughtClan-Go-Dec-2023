package main

import "fmt"

func main() {
	var no int = 100
	fmt.Println("Before incrementing, no :", no)
	increment(&no)
	fmt.Println("After incrementing, no :", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("Aftter swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) /* NO RETURN RESULT */ {
	*x++
}

func swap(x, y *int) /* NO RETURN RESULT */ {
	*x, *y = *y, *x
}
