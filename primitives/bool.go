package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var temp bool // default => false
	temp2 := true // default => float64

	fmt.Printf("temp (default value) = %v => Type: %T | Value: %v | Size: %d bits\n", temp, temp, temp, unsafe.Sizeof(temp))
	// temp (default value) = false => Type: bool | Value: false | Size: 1 bits

	fmt.Printf("temp2 = %v => Type: %T | Value: %v | Size: %d bits\n", temp2, temp2, temp2, unsafe.Sizeof(temp2))
	// temp2 = true => Type: bool | Value: true | Size: 1 bits

}
