package luhn

import (
	"strings"
	"unicode"
)

// Valid determines whether a string is valid per the Luhn formula.
// The Luhn algorithm is used to validate identification numbers like credit cards.
//
// Rules:
// - Strings of length 1 or less are not valid
// - Spaces are allowed and should be stripped
// - All other non-digit characters are disallowed
// - Double every second digit from the right
// - If doubling results in > 9, subtract 9
// - Sum all digits and check if divisible by 10
func Valid(id string) bool {
	// Remove all spaces
	cleaned := strings.ReplaceAll(id, " ", "")
	
	// Check if length is valid (must be > 1)
	if len(cleaned) <= 1 {
		return false
	}
	
	// Check that all characters are digits
	for _, char := range cleaned {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	
	// Convert string to slice of digits and process from right to left
	digits := make([]int, len(cleaned))
	for i, char := range cleaned {
		digits[i] = int(char - '0')
	}
	
	// Apply Luhn algorithm
	sum := 0
	isSecondDigit := false
	
	// Process digits from right to left
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		
		// Double every second digit from the right
		if isSecondDigit {
			digit *= 2
			// If doubling results in > 9, subtract 9
			if digit > 9 {
				digit -= 9
			}
		}
		
		sum += digit
		isSecondDigit = !isSecondDigit
	}
	
	// Valid if sum is divisible by 10
	return sum%10 == 0
}
