package main

import "fmt"

func main() {
	// var x interface{}
	var x any // any is an alias for interface{}
	x = 100
	x = "Nisi consequat commodo occaecat esse. Laboris voluptate et magna enim nostrud occaecat nisi reprehenderit irure incididunt cillum culpa minim do. Culpa et culpa aliquip sit fugiat labore id adipisicing nisi ad. Qui labore culpa dolore amet sit exercitation. Commodo reprehenderit proident qui duis qui. Duis culpa anim ad do magna et adipisicing culpa."
	x = 99.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	x = 100
	// x = "Proident excepteur eu aliqua ullamco aute aliqua. Dolor aliquip ut ad sunt ea cillum fugiat pariatur amet eu. Et Lorem in voluptate cillum voluptate reprehenderit duis deserunt."
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type switch
	// x = 100
	// x = "Nisi consequat commodo occaecat esse. Laboris voluptate et magna enim nostrud occaecat nisi reprehenderit irure incididunt cillum culpa minim do. Culpa et culpa aliquip sit fugiat labore id adipisicing nisi ad. Qui labore culpa dolore amet sit exercitation. Commodo reprehenderit proident qui duis qui. Duis culpa anim ad do magna et adipisicing culpa."
	// x = 99.99
	// x = true
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 5 =", val*5)
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	case float64:
		fmt.Println("x is a float64, 0.09% of x =", val*0.09)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	default:
		fmt.Println("x is of unknown type")
	}

}
