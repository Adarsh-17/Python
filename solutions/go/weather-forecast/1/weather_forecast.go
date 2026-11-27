//Package weather provides tools to get forecast.
package weather

//CurrentCondition stores the current condition.
var CurrentCondition string
//CurrentLocation stores the current location.
var CurrentLocation string

//Forecast returns weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
