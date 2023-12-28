package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating (indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating (range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// adding new items
	// nos = append(nos, 10)
	nos = append(nos, 10, 20, 30, 40)
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[2:5] =", nos[2:5])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:5] =", nos[:5])

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos = %v, nos2 = %v\n", nos, nos2)

	tens := nos[5:9]
	fmt.Println("tens[0] = ", tens[0])
	fmt.Println("tens[1] = ", tens[1])
	fmt.Println("tens[2] = ", tens[2])
	fmt.Println("tens[3] = ", tens[3])

	// fmt.Println("tens[4] = ", tens[4]) // => index out of range error panic

	sort(nos)
	fmt.Println(nos)

	// duplicating a slice
	dupNos := make([]int, len(nos))
	copy(dupNos, nos)
	fmt.Println(dupNos)
	nos[0] = 7777
	fmt.Println(nos)
	fmt.Println(dupNos)
}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
