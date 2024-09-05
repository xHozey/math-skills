package main

import (
	"fmt"
	"math"

	skill "math_skills/assests"
)

func main() {
	data := skill.GetData()
	averge := skill.Averge(data)
	median := skill.Median(data)
	variance := skill.Variance(data, averge)
	stdDev := math.Sqrt(variance)
	standardDev := math.Round(stdDev)
	fmt.Println("Averge:", averge)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", standardDev)
}
