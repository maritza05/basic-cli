package main

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

func TestDummy(t *testing.T) {

	result := 1 + 1
	if result != 2 {
		t.Errorf("expected '%d' but got '%d'", 2, result)
	}

	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})

	x := []float64{1.0, 1.1, 1.2, math.Pi}
	y := []float64{1.0, 1.1, 1.2, 3.14159265359} // Accurate enough to Pi

	if !cmp.Equal(x, y, opt) {
		t.Errorf("expected '%d' but got '%d'", 2, result)
	}
}
