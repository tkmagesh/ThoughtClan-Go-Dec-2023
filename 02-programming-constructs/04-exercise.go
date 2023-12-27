/* Print all the prime numbers between 2 to 100 */
package main

import (
	"fmt"
	"math"
)

func main() {
OUTER_LOOP:
	for no := 2; no <= 100; no++ {
		limit := int(math.Sqrt(float64(no)))
		for i := 2; i <= limit; i++ {
			if no%i == 0 {
				continue OUTER_LOOP
			}
		}
		fmt.Println("Prime no :", no)
	}
}
