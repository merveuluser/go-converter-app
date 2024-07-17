package helper

import (
	"math"
)

func RoundFloat(v float64) float64 {
	ratio := math.Pow(10, 2)
	return math.Round(v*ratio) / ratio
}
