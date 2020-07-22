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
	fmt.Printf("%v => Type: %T | Value: %v | Size: %d bits\n\n", temp, temp, temp, unsafe.Sizeof(temp))
	// 1 => Type: int | Value: 1 | Size: 8 bits

	fmt.Printf("uint(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, uint(temp), uint(temp), unsafe.Sizeof(uint(temp)))
	fmt.Printf("uint8(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint8(temp), uint8(temp), unsafe.Sizeof(uint8(temp)))
	fmt.Printf("uint16(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint16(temp), uint16(temp), unsafe.Sizeof(uint16(temp)))
	fmt.Printf("uint32(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, uint32(temp), uint32(temp), unsafe.Sizeof(uint32(temp)))
	fmt.Printf("uint64(%v) => Type: %T | Value: %v | Size: %d bytes\n\n", temp, uint64(temp), uint64(temp), unsafe.Sizeof(uint64(temp)))

	// uint(1) => Type: uint | Value: 1 | Size: 8 bits
	// uint8(1) => Type: uint8 | Value: 1 | Size: 1 bytes
	// uint16(1) => Type: uint16 | Value: 1 | Size: 2 bytes
	// uint32(1) => Type: uint32 | Value: 1 | Size: 4 bytes
	// uint64(1) => Type: uint64 | Value: 1 | Size: 8 bytes

	fmt.Printf("int(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, int(temp), int(temp), unsafe.Sizeof(int(temp)))
	fmt.Printf("int8(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, int8(temp), int8(temp), unsafe.Sizeof(int8(temp)))
	fmt.Printf("int16(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, int16(temp), int16(temp), unsafe.Sizeof(int16(temp)))
	fmt.Printf("int32(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, int32(temp), int32(temp), unsafe.Sizeof(int32(temp)))
	fmt.Printf("int64(%v) => Type: %T | Value: %v | Size: %d bytes\n", temp, int64(temp), int64(temp), unsafe.Sizeof(int64(temp)))

	// int(1) => Type: int | Value: 1 | Size: 8 bits
	// int8(1) => Type: int8 | Value: 1 | Size: 1 bytes
	// int16(1) => Type: int16 | Value: 1 | Size: 2 bytes
	// int32(1) => Type: int32 | Value: 1 | Size: 4 bytes
	// int64(1) => Type: int64 | Value: 1 | Size: 8 bytes
}
