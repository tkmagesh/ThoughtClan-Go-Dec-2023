package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()
	sayHi = func() {
		fmt.Println("Hello there!")
	}
	sayHi()

	/* 1 parameter, no return results */
	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")
}

/* 2 parameters, no return results */
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

/* 2 parameters, 1 return result */
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

/* 2 parameters, 2 return results  */
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
