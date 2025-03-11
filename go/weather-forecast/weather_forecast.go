/* Package weather returns weather conditions around Goblinocus. */
package weather

// CurrentCondition holds the value of the current weather condition.
var CurrentCondition string

// CurrentLocation public var that holds the location being examined.
var CurrentLocation string

// Forecast returns the current weather condition for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
