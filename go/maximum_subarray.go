package main

func MaxSubarray(a []int) []int {
	currentSumStart := 0
	currentSum := 0
	maxSumStart := 0
	maxSumEnd := 0
	maxSum := 0

	for currentSumEnd := 0; currentSumEnd < len(a); currentSumEnd++ {
		currentSum += a[currentSumEnd]

		if currentSum > maxSum {
			maxSum = currentSum
			maxSumStart = currentSumStart
			maxSumEnd = currentSumEnd
		}

		if currentSum < 0 {
			currentSumStart = currentSumEnd + 1
			currentSum = 0
		}
	}
	return a[maxSumStart : maxSumEnd + 1]
}
