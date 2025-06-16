package raindrops

// Convert returns a string based on the number's divisibility
func Convert(number int) string {
    result := ""

    if number%3 == 0 {
        result += "Pling"
    }
    if number%5 == 0 {
        result += "Plang"
    }
    if number%7 == 0 {
        result += "Plong"
    }

    if result == "" {
        return intToString(number)
    }

    return result
}

// intToString converts an integer to a string without using strconv
func intToString(n int) string {
    if n == 0 {
        return "0"
    }

    isNegative := n < 0
    if isNegative {
        n = -n
    }

    digits := []byte{}
    for n > 0 {
        digit := n % 10
        digits = append([]byte{byte('0' + digit)}, digits...)
        n /= 10
    }

    if isNegative {
        digits = append([]byte{'-'}, digits...)
    }

    return string(digits)
}