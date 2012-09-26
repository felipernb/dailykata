package sqrt

import (
	"testing"
	"fmt"
	"math"
)

func assertEquals(a, b, delta float64, t *testing.T) {
	if math.Abs(a-b)/b > delta {
		t.Fail()
		fmt.Printf("FAIL: %d is not close enough to %d", b, a)
	}
}

func TestSqrt(t *testing.T) {
	for i := 0.0; i < 1000; i++ {
		// Tests comparing with native library's method
		assertEquals(math.Sqrt(i), Sqrt(i), DELTA, t)
	}
}
