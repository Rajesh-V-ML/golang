package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Errorf("Expected 2 + 2 to equal 4")
	}
}
