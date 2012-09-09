package roman_to_decimal

func RomanToDecimal(roman string) int {
	table := map[uint8] int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := 0
	posteriorSymbol := 0
	// traverses the string backwards, if the posterior symbol
	// is greater than the current one, decrement (e.g.: IV => 5 - 1),
	// otherwise, sum
	for i := len(roman) - 1; i >= 0; i-- {
		currentSymbol := table[roman[i]]
		if (posteriorSymbol > currentSymbol) {
			n -= currentSymbol
		} else {
			n += currentSymbol
		}
		posteriorSymbol = currentSymbol
	}
	return n
}
