package main

import (
	"fmt"
)

func main() {
	customerName := "Judy"
	welcome := WelcomeMessage(customerName)
	fmt.Println("Welcome Message:")
	fmt.Println(welcome)
	fmt.Println()

	bordered := AddBorder("Welcome!", 10)
	fmt.Println("Message with Border:")
	fmt.Println(bordered)
	fmt.Println()

	oldMessage := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	cleaned := CleanUpMessage(oldMessage)
	fmt.Println("Cleaned Marketing Message:")
	fmt.Println(cleaned)
}