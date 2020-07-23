// Golang Maps is a collection of unordered pairs of key-value
// - Keys are reference to a hash table
// - Due to its reference type it is inexpensive to pass, for example, for a 64-bit machine it takes 8 bytes and for a 32-bit machine, it takes 4 bytes.
// - The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.

// Syntax
// map[Key_Type]Value_Type{}
// var mymap map[int]string

// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}
// var mymap map[int]string{1:"a",2:"b"}


package main

import (
	"fmt"
)

func main() {
	var map_1 map[int]int

	// Checking if the map is nil or not 
	if map_1 == nil {   // True

	fmt.Println("True") 
	} else { 

	fmt.Println("False") 
	} 

	// Creating and initializing a map 
	// Using shorthand declaration and 
	// using map literals 
	map_2 := map[int]string{ 

	    90: "Dog", 
	    91: "Cat", 
	    92: "Cow", 
	    93: "Bird", 
	    94: "Rabbit", 
	} 

	fmt.Println("Map-2: ", map_2) // Map-2:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
	
	// Iterating map using for rang loop 
	for id, pet := range map_2 { 
		fmt.Println(id, pet) 
	} 
	// 90 Dog
	// 91 Cat
	// 92 Cow
	// 93 Bird
	// 94 Rabbit
	
	// Checking the key is available 
	// or not in the map_2 map 
	pet_name, ok := map_2[90] 
	fmt.Println("\nKey present or not:", ok) 
	fmt.Println("Value:", pet_name) 

	// Using blank identifier 
	_, ok1 := map_2[92] 
	fmt.Println("\nKey present or not:", ok1) 
	
	// Creating a map using make() function 
	var My_map = make(map[float64]string)  // map[]
	My_map[1.3] = "Rohit"
	My_map[1.5] = "Sumit"
	fmt.Println(My_map) // map[1.3:Rohit 1.5:Sumit]
	
	
	
	
	
	
}
