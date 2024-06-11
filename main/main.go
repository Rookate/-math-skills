package main

import (
	"fmt"
	"math"
	"os"
	"skills"
	"strconv"
)

func main() {
	data := skills.Read_File(os.Args[1])

	average := skills.Average(data)
	fmt.Println("Average:", math.Round(average))

	median := skills.Median(data)
	fmt.Println("Median:", math.Round(median))

	variance := skills.Variance(data)
	conv := strconv.FormatInt(int64(math.Round(variance)), 10)
	fmt.Println("Variance:", conv)

	standard_variation := skills.Standard(data)
	fmt.Println("Standard Variation:", math.Round(standard_variation))
}
