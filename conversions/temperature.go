// Package conversions contain main conversion functions.
package conversions

import (
	"github.com/merveuluser/go-converter-app/constants"
)

// CelsiusToFahrenheit converts given Celsius value to its Fahrenheit value.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*constants.CToFFactor + constants.FreezingPointF
}

// FahrenheitToCelsius converts given Fahrenheit value to its Celsius value.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - constants.FreezingPointF) * constants.FToCFactor
}

// CelsiusToKelvin converts given Celsius value to its Kelvin value.
func CelsiusToKelvin(celsius float64) float64 {
	return celsius + constants.CToK
}

// KelvinToCelsius converts given Kelvin value to its Celsius value.
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin + constants.KToC
}

// FahrenheitToKelvin converts given Fahrenheit value to its Kelvin value.
func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-constants.FreezingPointF)*constants.FToCFactor + constants.CToK
}

// KelvinToFahrenheit converts given Kelvin value to its Fahrenheit value.
func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin+constants.KToC)*constants.CToFFactor + constants.FreezingPointF
}
