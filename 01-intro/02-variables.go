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
}
