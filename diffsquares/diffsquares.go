package diffsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers
// Formula: (1 + 2 + ... + n)²
// The sum of first n natural numbers is n(n+1)/2
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers
// Formula: 1² + 2² + ... + n²
// The sum of squares formula is n(n+1)(2n+1)/6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the square of the sum
// and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
