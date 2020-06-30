/*		Package Comments Format

package go_packages  // no under_scores or mixedCaps allowed
*/
package gopackages

import "fmt"

// Exported function :
// Function name with capital intitial
// This will be avaible to other packages
// Exported Function comments are written with "//"
func Exported(a, b int) {
	fmt.Println("Exported : I will be exported to other packages.")
}

// Functions are lower Camel case
func unExported(a, b int) {
	fmt.Println("unExported : I will be not be availabe to other packages.")
}

// Functions are lower Camel case
func returnTypeInt(a, b int) int {
	fmt.Println("I need a return value")
	return 0
}
