package cars

// CalculateWorkingCarsPerHour calculates the number of working cars produced per hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates the number of working cars produced per minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost calculates the cost of production based on the number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individuals := carsCount % 10
	return uint(groupsOfTen*95000 + individuals*10000)
}
