/*
Build an interactive calculator
- Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit
- if input is 5
	then exit
- if input > 5 OR < 1
	print "invalid input"
	display the menu again
- if input >=1 OR <= 4
	accept 2 operands
	perform the corresponding operation
	print the result
	display the menu again

Assumption:
	The user will always enter valid data
*/

package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int

	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
