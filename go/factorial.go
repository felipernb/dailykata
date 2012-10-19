package main

func factorial(n int) int64 {
	fact := int64(1)
	for i := int64(n); i > 0; i-- { fact *= i }
	return fact
}
