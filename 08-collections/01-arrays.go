package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// use :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating (indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating (range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // creating a copy of the array
	nos2[0] = 9999
	fmt.Println("nos :", nos)
	fmt.Println("nos2 :", nos2)

	// array pointers
	var nosPtr *[5]int
	nosPtr = &nos
	fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0])

	sort(&nos)
	fmt.Println("After sorting, nos : ", nos)
}

func sort(list *[5]int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
