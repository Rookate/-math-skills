package skills

import (
	"math"
)

func Standard(data []int) float64 {
	variance := Variance(data)

	deviation := math.Sqrt(variance)

	return deviation
}
