package sorting

import (
    "math/rand"
)

func Quicksort(a []int) []int {
	if (len(a) <= 1) {
		return a
	}
	p := partition(a)
    Quicksort(a[0:p])
	Quicksort(a[p+1:])
	return a
}

func partition(a []int) int {
    p := rand.Intn(len(a)) // Random pivot
    a[0], a[p] = a[p], a[0]
	pivot := a[0]
	i := 1
	for j := i; j < len(a); j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[0], a[i-1] = a[i-1], a[0]
	return i-1
}
