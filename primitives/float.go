// float32 float64
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

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

	var f float64 = 123.123456

	fmt.Printf("%0.2f \n", f)
	fmt.Printf("%0.2f \n", Round(f)) // round half
}
