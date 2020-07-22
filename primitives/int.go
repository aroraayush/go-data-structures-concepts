// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// fmt.Sprintf
	// Format a string without printing it
	s := fmt.Sprintf("Sprintf(a + b) = %d + %d \n", 4, 5)
	fmt.Println(s)

	// Integer examples

	var temp int = 1

	// format a string without printing it
	// fmt.Printf takes type string in argument
	fmt.Printf("%v 	=> Type: %T | Value: %v | Size: %d bits\n", temp, temp, temp, unsafe.Sizeof(temp))
	fmt.Printf("uint(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, uint(temp), uint(temp), unsafe.Sizeof(uint(temp)))
	fmt.Printf("uint8(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint8(temp), uint8(temp), unsafe.Sizeof(uint8(temp)))
	fmt.Printf("uint16(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint16(temp), uint16(temp), unsafe.Sizeof(uint16(temp)))
	fmt.Printf("uint32(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint32(temp), uint32(temp), unsafe.Sizeof(uint32(temp)))
	fmt.Printf("uint64(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint64(temp), uint64(temp), unsafe.Sizeof(uint64(temp)))
}
