package main

import "fmt"

// CANNOT use ":=" at package scope
// app_version := "1.0"

// var app_version string = "1.0"
var app_version = "1.0"

func main() {
	/*
		var username string
		username = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", username)
	*/

	/*
		var username string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", username)
	*/

	// type inference
	/*
		var username = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", username)
	*/

	username := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", username)

	/* unused variable */
	/*
		var x int
		x = 100
	*/
	// fmt.Println(x)

	/* handling multiple variables */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y int    = 100, 200
			str  string = "Sum of %d and %d is %d\n"
		)
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var x = 100
		var y = 200
		var str = "Sum of %d and %d is %d\n"
		var result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x      = 100
			y      = 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}
