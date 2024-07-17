// Package conversions contain main conversion functions.
package conversions

import "github.com/merveuluser/go-converter-app/constants"

func GramsToOunces(meters float64) float64 {
	return meters * constants.GToOz
}

func KilogramsToPounds(kilometers float64) float64 {
	return kilometers * constants.KgToLbs
}
