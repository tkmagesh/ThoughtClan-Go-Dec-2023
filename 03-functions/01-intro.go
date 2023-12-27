package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")

	msg := getGreetMsg("Suresh", "Kumar")
	fmt.Print(msg)

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// print only quotient
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/
}

/* no parameters, no return results */
func sayHi() {
	fmt.Println("Hi there!")
}

/* 1 parameter, no return results */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
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
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/* 2 parameters, 2 return results (named results)  */
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
