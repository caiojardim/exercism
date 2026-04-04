// Package weather provides a function that return the weather conditions.
package weather

var (
	//CurrentCondition reprensets the current weather condition.
	CurrentCondition string
	//CurrentLocation reprensets the locantion of the wheather condition.
	CurrentLocation string
)

// Forecast return an String with the current locantion and the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
