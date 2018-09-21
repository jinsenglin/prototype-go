package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	if result := Sqrt(1); result != 1 {
		t.Errorf("Input 1 | Expected Output 1 | Returned Output %v", result)
	}

	if result := Sqrt(4); result != 2 {
		t.Errorf("Input 4 | Expected Output 2 | Returned Output %v", result)
	}
}
