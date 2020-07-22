// float32 float64
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	temp := 1.1 // default => float64

	fmt.Printf("%v => Type: %T | Value: %v | Size: %d bits\n\n", temp, temp, temp, unsafe.Sizeof(temp))
	// 1.1 => Type: float64 | Value: 1.1 | Size: 8 bits

	// float32
	fmt.Printf("float32(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, float32(temp), float32(temp), unsafe.Sizeof(float32(temp)))
	// float32(1.1) => Type: float32 | Value: 1.1 | Size: 4 bits

	// float64
	fmt.Printf("float64(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, float64(temp), float64(temp), unsafe.Sizeof(float64(temp)))
	// float64(1.1) => Type: float64 | Value: 1.1 | Size: 8 bits
}
