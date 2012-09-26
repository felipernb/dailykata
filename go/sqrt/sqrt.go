package sqrt

import (
	"math"
)
const (
	DELTA = 10E-10
)

func Sqrt(n float64) float64 {
	x := n/2
	for math.Abs(x*x - n)/n > DELTA {
		x = (x+n/x)/2
	}

	// Optimization for perfect squares
	roundedX := math.Floor(x+0.5)
	if roundedX * roundedX == n {
		x = roundedX
	}
	return x
}
