// Package twofer implements a sharing message generator.
// It follows the "one for X, one for me" pattern, where X can be a name or "you" if no name is provided.
//
// Example usage:
//
//	twofer.ShareWith("Alice") // returns "One for Alice, one for me."
//	twofer.ShareWith("")      // returns "One for you, one for me."
package twofer

// ShareWith returns a message indicating sharing something with the given name.
// If the name is empty, it defaults to "you".
//
// Examples:
//
//	ShareWith("Alice") // returns "One for Alice, one for me."
//	ShareWith("")      // returns "One for you, one for me."
//	ShareWith("Bob")   // returns "One for Bob, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
