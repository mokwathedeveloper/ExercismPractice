package main

// ToUpperCase converts all lowercase letters in the input string to uppercase.
func ToUpperCase(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r -= 'a' - 'A'
		}
		result += string(r)
	}
	return result
}

// RepeatString returns a new string consisting of the input string repeated 'count' times.
func RepeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// ReplaceAllChars replaces all occurrences of a substring 'old' with a substring 'new'.
func ReplaceAllChars(s, old, new string) string {
	result := ""
	i := 0
	for i < len(s) {
		if len(s[i:]) >= len(old) && s[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}

// TrimSpaces removes all leading and trailing whitespace characters.
func TrimSpaces(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && (s[start] == ' ' || s[start] == '\n' || s[start] == '\t') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\n' || s[end] == '\t') {
		end--
	}
	return s[start : end+1]
}

// WelcomeMessage generates a personalized welcome message.
func WelcomeMessage(name string) string {
	return "Welcome to the Tech Palace, " + ToUpperCase(name)
}

// AddBorder wraps the message in a star border.
func AddBorder(message string, numStars int) string {
	border := RepeatString("*", numStars)
	return border + "\n" + message + "\n" + border
}

// CleanUpMessage removes stars and trims the message.
func CleanUpMessage(message string) string {
	clean := ReplaceAllChars(message, "*", "")
	return TrimSpaces(clean)
}
