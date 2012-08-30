package main

import (
    "math/rand"
)

func RSelect(a []int, i int) int {
    n := len(a)
    if n == 1 {
        return a[0]
    }
    p := rand.Intn(n)
    a[0], a[p] = a[p], a[0] //put pivot as the first
    j := 0
    for k := 1; k < n; k++ {
        if a[k] < a[0] {
            a[j+1], a[k] = a[k], a[j+1]
            j++
        }
    }
    a[0], a[j] = a[j], a[0]
    if j == i-1 {
        return a[j]
    } else if j > i-1 {
        return RSelect(a[0:j], i)
    }
    return RSelect(a[j:], i-j)
}
