package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remainingTime := OvenTime - actualMinutesInOven
	if remainingTime < 0 {
		return 0
	} else {
		return remainingTime
	}
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	PreparationTime := numberOfLayers * 2
	return PreparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven
	return elapsedTime
}
