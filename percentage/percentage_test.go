package percentage

import (
	"testing"
)

func TestCanParsePercentage(t *testing.T) {
	var want float32 = 0.05
	// TODO: Add table tests to check different strings
	got, err := ParsePercentToFloat32("5%")
	if err != nil {
		t.Fail()
	}
	if want != got {
		t.Fail()
	}
	want = 1.00
	got, err = ParsePercentToFloat32("    100%     ")
	if err != nil {
		t.Fail()
	}
	if want != got {
		t.Fail()
	}
}
