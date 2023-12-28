package main

import (
	"fmt"
)

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 4
		productRanks["pencil"] = 1
		productRanks["marker"] = 2
	*/

	// var productRanks map[string]int = map[string]int{"pen": 4, "pencil": 1, "marker": 2}
	var productRanks map[string]int = map[string]int{
		"pen":    4,
		"pencil": 1,
		"marker": 2,
	}
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) : ", len(productRanks))

	fmt.Println("Iterating using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// keyToRemove := "pen"
	keyToRemove := "notebook" //non existing key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key %q\n", keyToRemove)
	fmt.Println(productRanks)

	// check for the existence of the key
	// keyToCheck := "marker"
	keyToCheck := "notebook" //non existing key
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

}
