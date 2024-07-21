package conversions

type ConversionFunc func(float64) float64

var ConversionMap = map[string]ConversionFunc{
	"meterstoinches":      MetersToInches,
	"kilometerstomiles":   KilometersToMiles,
	"gramstoounces":       GramsToOunces,
	"kilogramstopounds":   KilogramsToPounds,
	"celsiustofahrenheit": CelsiusToFahrenheit,
	"fahrenheittocelsius": FahrenheitToCelsius,
	"celsiustokelvin":     CelsiusToKelvin,
	"kelvintocelsius":     KelvinToCelsius,
	"fahrenheittokelvin":  FahrenheitToKelvin,
	"kelvintofahrenheit":  KelvinToFahrenheit,
}
