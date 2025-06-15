// Package twofer provides a function to generate a sharing phrase
// for the "two-fer" cookie sharing scenario.
package twofer

// ShareWith returns a string saying who the cookie is for.
// If the name is empty, it defaults to "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
