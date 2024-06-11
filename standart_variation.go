package skills

import (
	"math"
)

func Standart(data []int) float64 {
	variance := Variance(data)

	deviation := math.Sqrt(variance)

	return deviation
}
