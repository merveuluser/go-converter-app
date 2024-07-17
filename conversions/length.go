// Package conversions contain main conversion functions.
package conversions

import (
	"github.com/merveuluser/go-converter-app/constants"
)

func MetersToInches(meters float64) float64 {
	return meters * constants.MToIn
}

func KilometersToMiles(kilometers float64) float64 {
	return kilometers * constants.KmToMi
}
