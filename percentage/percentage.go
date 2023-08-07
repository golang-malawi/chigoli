package percentage

import (
	"fmt"
	"strconv"
	"strings"
)

// ParsePercentToFloat32 converts a human readable percent spec to a float.
// "5%" -> 0.05,
func ParsePercentToString(spec string) (string, error) {
	p, err := ParsePercentToFloat32(spec)
	if err != nil {
		return "0.0", err
	}

	return fmt.Sprintf("%3f", p), nil
}

func MustParsePercentToString(spec string) string {
	p, err := ParsePercentToFloat32(spec)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%3f", p)
}

func ParsePercentToFloat32(spec string) (float32, error) {
	n := strings.TrimSpace(spec)
	if !strings.HasSuffix(n, "%") {
		return 0.0, fmt.Errorf("string does not have percentage sign")
	}

	percentageFull, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(n, "%")), 32)
	if err != nil {
		return 0.0, err
	}
	return float32(percentageFull / 100.00), nil
}

func MustParsePercentToFloat32(spec string) float32 {
	v, err := ParsePercentToFloat32(spec)
	if err != nil {
		panic(err)
	}
	return v
}
