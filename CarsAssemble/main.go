package main

import (
	"fmt"
	"carsassemble/cars"
)

func main() {
	productionRate := 120
	successRate := 85.0
	carsProduced := 117

	workingCarsHour := cars.CalculateWorkingCarsPerHour(productionRate, successRate)
	workingCarsMinute := cars.CalculateWorkingCarsPerMinute(productionRate, successRate)
	cost := cars.CalculateCost(carsProduced)

	fmt.Printf("Working cars produced per hour: %.2f\n", workingCarsHour)
	fmt.Printf("Working cars produced per minute: %d\n", workingCarsMinute)
	fmt.Printf("Cost to produce %d cars: %d\n", carsProduced, cost)
}
