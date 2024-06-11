package skills

import "sort"

func Median(numbers []int) float64 {
	// Trier les nombres
	sort.Ints(numbers)

	n := len(numbers)
	if n == 0 {
		panic("La liste ne doit pas être vide")
	}

	// Si le nombre d'éléments est impair
	if n%2 != 0 {
		return float64(numbers[n/2])
	}

	// Si le nombre d'éléments est pair
	mid1 := numbers[(n/2)-1]
	mid2 := numbers[n/2]
	return float64(mid1+mid2) / 2
}
