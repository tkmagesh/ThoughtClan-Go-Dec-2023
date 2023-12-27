package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply result :", x*y)
	}, 100, 200)
}

func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started...")
	op(x, y)
	log.Println("Operation completed!")
}

func logAdd(x, y int) {
	log.Println("Operation started...")
	add(x, y)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started...")
	subtract(x, y)
	log.Println("Operation completed!")
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
