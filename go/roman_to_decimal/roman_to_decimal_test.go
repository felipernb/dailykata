package roman_to_decimal

import (
	"testing"
	"fmt"
)

func assertEquals(a, b int, t *testing.T) {
	if a != b {
		t.Fail()
		fmt.Printf("FAIL: Expected %d but got %d\n", a, b)
	}
}

func TestConversion(t *testing.T) {
	assertEquals(1, RomanToDecimal("I"), t)
	assertEquals(2, RomanToDecimal("II"), t)
	assertEquals(3, RomanToDecimal("III"), t)
	assertEquals(4, RomanToDecimal("IV"), t)
	assertEquals(5, RomanToDecimal("V"), t)
	assertEquals(6, RomanToDecimal("VI"), t)
	assertEquals(7, RomanToDecimal("VII"), t)
	assertEquals(10, RomanToDecimal("X"), t)
	assertEquals(20, RomanToDecimal("XX"), t)
	assertEquals(40, RomanToDecimal("XL"), t)
	assertEquals(50, RomanToDecimal("L"), t)
	assertEquals(60, RomanToDecimal("LX"), t)
	assertEquals(70, RomanToDecimal("LXX"), t)
	assertEquals(90, RomanToDecimal("XC"), t)
	assertEquals(100, RomanToDecimal("C"), t)

}

