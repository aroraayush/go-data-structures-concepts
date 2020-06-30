package main

import "fmt"

// Once an array is defined, elements cannot be added or removed from the array
// Go is all about types, we define the array with the type of element it

// Package scope variables
var integerArray [3]int // Variables are in camelCase
var stringArray [2]string

func main() {

	// fmt.Println("Hello, world.")
	integerArray[0] = 10
	integerArray[1] = 20
	integerArray[2] = 30

	stringArray[0] = "first"
	stringArray[1] = "second"

	//Function scope array
	integerArray2 := [5]int{10, 20, 30, 40, 50}
	stringArray2 := [4]string{"first", "second", "third", "fourth"}

	fmt.Println("This is the integer array: ", integerArray)
	fmt.Println("This is the string array: ", stringArray) // Quotes will not be printed

	fmt.Println("This is the integer array2: ", integerArray2)
	fmt.Println("This is the string array2: ", stringArray2)
}
