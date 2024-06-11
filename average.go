package skills

func Average(data []int) float64 {
	n := len(data)
	sum := 0

	for _, i := range data {
		sum += i
	}
	average := float64(sum) / float64(n)

	return average
}
