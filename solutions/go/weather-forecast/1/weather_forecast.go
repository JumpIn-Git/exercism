//Package weather, stores weather info.
package weather

//CurrentCondition is a description.
var CurrentCondition string
//CurrentLocation is the current city.
var CurrentLocation string

//Forecast returns a message and saves info.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
