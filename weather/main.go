package main

import (
	"fmt"
)

func main() {
	report := Forecast("Goblintown", "sunny with a chance of fireballs")
	fmt.Println(report)

	// Optional: access current state
	fmt.Println("Stored Location:", CurrentLocation)
	fmt.Println("Stored Condition:", CurrentCondition)
}
