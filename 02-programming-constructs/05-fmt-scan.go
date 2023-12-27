package main

import "fmt"

func main() {
	/*
		var username string
		fmt.Println("Enter your name :")
		// username = "Magesh"
		fmt.Scanln(&username)
		fmt.Printf("Hi %s, Have a nice day!\n", username)
	*/

	var n1, n2 int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1 + n2)
}
