package skills

import (
	"math"
)

func Variance(data []int) float64 {
	number := Average(data)
	var sd float64

	for i := 0; i < len(data); i++ {
		sd += math.Pow(float64(data[i])-number, 2)
	}

	return sd / float64(len(data)) // Divisez la somme des carrés par le nombre total d'éléments
}
