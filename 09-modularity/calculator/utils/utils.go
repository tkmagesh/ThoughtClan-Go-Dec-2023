package utils

import "fmt"

func init() {
	fmt.Println("utils package initialized")
}

func IsEven(no int) bool {
	return no%2 == 0
}
