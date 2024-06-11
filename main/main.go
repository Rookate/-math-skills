package main

import (
	"fmt"
	"math"
	"os"
	"skills"
)

func main() {
	data := skills.Read_File(os.Args[1])

	average := skills.Average(data)
	fmt.Println("Average:", math.Round(average))

	median := skills.Median(data)
	fmt.Println("Median:", math.Round(median))
}
