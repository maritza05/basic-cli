package main

import (
	"testing"
)

func TestDummy(t *testing.T) {

	result := 1 + 1
	if result != 2 {
		t.Errorf("expected '%d' but got '%d'", 2, result)
	}
}
