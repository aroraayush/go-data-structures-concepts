// float32 float64
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	temp := 1
	fmt.Printf("float32(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, float32(temp), float32(temp), unsafe.Sizeof(float32(temp)))
	fmt.Printf("float64(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, float64(temp), float64(temp), unsafe.Sizeof(float64(temp)))
}
