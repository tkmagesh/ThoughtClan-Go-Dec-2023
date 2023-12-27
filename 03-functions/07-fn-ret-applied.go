package main

import (
	"fmt"
	"log"
	"time"
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

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		}, 100, 200)
	*/

	// composing "log" functionality
	/*
		var logAdd func(int, int)
		logAdd = getLogOperation(add)
		logAdd(100, 200)

		var logSubtract func(int, int)
		logSubtract = getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// composing "profile" functionality
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)

		profiledMultiply := getProfiledOperation(func(x, y int) {
			fmt.Println("Multiply Result :", x*y)
		})
		profiledMultiply(100, 200)
	*/

	// composing "log" & "profile" functionality
	logAdd := getLogOperation(add)
	profiledLogAdd := getProfiledOperation(logAdd)
	profiledLogAdd(100, 200)

	getProfiledOperation(getLogOperation(subtract))(100, 200)
}

func getProfiledOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("[Profile] Elapsed :", elapsed)
	}
}

func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started...")
		op(x, y)
		log.Println("Operation completed!")
	}
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
