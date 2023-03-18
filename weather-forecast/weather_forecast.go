// Package weather: forecast weather.
package weather

// CurrentCondition: weather condition.
var CurrentCondition string

// CurrentLocation: city location.
var CurrentLocation string

// Forecast: forecast weather for given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
