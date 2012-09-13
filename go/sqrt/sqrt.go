package sqrt

import (
	"math"
)
const (
	DELTA = 10E-10
)

func Sqrt(n float64) float64 {
	lower := 0.0
	upper := n
	x := n/2
	for math.Abs(x*x - n) > DELTA {
		if x*x > n {
			upper = x
		} else {
			lower = x
		}
		x = (upper - lower)/2 + lower
	}

	// Optimization for perfect squares
	roundedX := math.Floor(x+0.5)
	if roundedX * roundedX == n {
		x = roundedX
	}
	return x
}
