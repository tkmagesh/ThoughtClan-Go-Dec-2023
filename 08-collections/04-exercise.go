/* Write a function "genPrimes" that generates the prime numbers between the given start & end

Use the "genPrimes" function in "main" function to generate the prime numbers between 2 and 100 and print
them in the main() function
*/

package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, no := range primes {
		fmt.Println(no)
	}
}

func genPrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
