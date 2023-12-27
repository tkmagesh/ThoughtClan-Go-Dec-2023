/* for type conversion - use the type name like a function */
package main

import "fmt"

func main() {
	var i int32 = 100
	var f float64
	f = float64(i)
	fmt.Println(f)
}
