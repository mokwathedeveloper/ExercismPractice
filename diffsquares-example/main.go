package main

import (
	"diffsquares"
	"fmt"
	"strings"
)

func main() {
	// Test with different values
	testValues := []int{1, 5, 10, 100}

	fmt.Println("Difference of Squares Calculator")
	fmt.Println("================================")

	for _, n := range testValues {
		fmt.Printf("\nFor the first %d natural numbers:\n", n)

		squareOfSum := diffsquares.SquareOfSum(n)
		sumOfSquares := diffsquares.SumOfSquares(n)
		difference := diffsquares.Difference(n)

		fmt.Printf("  Square of sum: (%d)² = %d\n", sumOfFirstN(n), squareOfSum)
		fmt.Printf("  Sum of squares: %d\n", sumOfSquares)
		fmt.Printf("  Difference: %d - %d = %d\n", squareOfSum, sumOfSquares, difference)
	}

	// Show the specific example from the problem
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Problem Example (n=10):")
	fmt.Printf("The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = %d\n", diffsquares.SquareOfSum(10))
	fmt.Printf("The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = %d\n", diffsquares.SumOfSquares(10))
	fmt.Printf("Hence the difference is %d - %d = %d\n", diffsquares.SquareOfSum(10), diffsquares.SumOfSquares(10), diffsquares.Difference(10))
}

// Helper function to calculate sum of first n natural numbers
func sumOfFirstN(n int) int {
	return n * (n + 1) / 2
}
