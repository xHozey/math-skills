package math_skills

import (
	"math"
	"sort"
)

func Averge(nb []float64) float64 {
	var result float64
	for i := 0; i < len(nb); i++ {
		result += float64(nb[i])
	}
	result /= float64(len(nb))
	n := math.Round(result)
	return n
}

func Median(nb []float64) float64 {
	var result float64
	sort.Float64s(nb)
	n := len(nb)
	if n%2 == 0 {

		mid1 := n / 2
		mid1 -= 1

		result = (nb[mid1] + nb[mid1+1]) / 2
	} else {

		mid := n / 2
		result = nb[mid]
	}

	return math.Round(result)
}

func Variance(nb []float64, average float64) float64 {
	var vSum float64

	for i := 0; i < len(nb); i++ {
		diff := nb[i] - average
		vSum += diff * diff
	}

	return math.Round(vSum / float64(len(nb)))
}
