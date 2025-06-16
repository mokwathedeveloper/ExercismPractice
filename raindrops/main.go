package main

import (
    "fmt"
    "raindrops/raindrops"
)

func main() {
    testNumbers := []int{28, 30, 34, 105, 3, 5, 7, 1, 0, -21}
    for _, num := range testNumbers {
        fmt.Printf("Convert(%d) = %s\n", num, raindrops.Convert(num))
    }
}