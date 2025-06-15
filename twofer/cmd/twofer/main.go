package main

import (
	"fmt"
	"twofer"
)

func main() {
	// Examples of using the ShareWith function
	fmt.Println(twofer.ShareWith("Alice"))
	fmt.Println(twofer.ShareWith("Bohdan"))
	fmt.Println(twofer.ShareWith(""))
}
