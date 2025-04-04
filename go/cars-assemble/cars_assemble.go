package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
	leftover := carsCount % 10
	cost := uint(groups) * 95000
	if leftover > 0 {
		cost += uint(leftover) * 10000
	}
	return cost
}
