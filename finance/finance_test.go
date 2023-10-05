package finance

import (
	"testing"
)

func TestCalculateCompoundInterest(t *testing.T) {
	var want float64 = 14400
	got := CalculateCompoundInterest(10000, 20, 2)

	if want != got {
		t.Fail()
	}
}

func TestCalculateROI(t *testing.T) {
	var want float64 = 20
	got := CalculateROI(10000, 12000)

	if want != got {
		t.Fail()
	}
}