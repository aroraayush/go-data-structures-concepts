// complex64 complex128
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	temp := 1 + 2i // default => complex128

	fmt.Printf("temp = %v => Type: %T | Value: %v | Size: %d bits\n", temp, temp, temp, unsafe.Sizeof(temp))
	// temp = (1+2i) => Type: complex128 | Value: (1+2i) | Size: 16 bits

	// complex64
	fmt.Printf("complex64(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, complex64(temp), complex64(temp), unsafe.Sizeof(complex64(temp)))
	// complex64((1+2i)) => Type: complex64 | Value: (1+2i) | Size: 8 bits

	// complex128
	fmt.Printf("complex128(%v) => Type: %T | Value: %v | Size: %d bits\n", temp, complex128(temp), complex128(temp), unsafe.Sizeof(complex128(temp)))
	// complex128((1+2i)) => Type: complex128 | Value: (1+2i) | Size: 16 bits
}
