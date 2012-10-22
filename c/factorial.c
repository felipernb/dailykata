long recurFact(int n) {
	if (n == 0) return 1L;
	return n * recurFact(n-1);
}

long iterFact(int n) {
	long fact = 1;
	for (; n > 1; n--) fact *= n;
	return fact;
}
