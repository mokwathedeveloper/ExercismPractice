package luhn

import "fmt"

// Example demonstrates how to use the Valid function
func Example() {
	fmt.Println("Luhn Algorithm Implementation")
	fmt.Println("============================")

	// Example usage:
	fmt.Println("Valid('4539 3195 0343 6467'):", Valid("4539 3195 0343 6467"))
	fmt.Println("Valid('8273 1232 7352 0569'):", Valid("8273 1232 7352 0569"))
	fmt.Println("Valid('1'):", Valid("1"))
}
