package main



import (
	"fmt"
	"ScrabbleScore/scrabble"
)

func main() {
	var word string
	fmt.Print("Enter a word to calculate its Scrabble score: ")
	fmt.Scanln(&word)

	score := scrabble.Score(word)
	fmt.Printf("The Scrabble score for '%s' is: %d\n", word, score)
}
