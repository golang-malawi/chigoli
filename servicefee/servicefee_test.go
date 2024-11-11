package servicefee_test

import (
	"testing"

	"github.com/golang-malawi/chigoli/servicefee"
	"github.com/stretchr/testify/assert"
)

func TestFixedFeesTable(t *testing.T) {
	serviceFees := servicefee.NewFixedFee(servicefee.FeeExpressions{
		"x <= 5000":               30.00,
		"x >= 5000 && x <= 10000": 20.00,
	})

	fee, err := serviceFees.CalculateFee(10_000)
	assert.Nil(t, err)
	assert.Equal(t, 20.00, fee)

	total, fee, err := serviceFees.CalculateTotalAndFee(10_000)
	assert.Nil(t, err)
	assert.Equal(t, 20.00, fee)
	assert.Equal(t, 10_020.00, total)
}

func TestPercentageFeesTable(t *testing.T) {
	serviceFees := servicefee.NewPercentageFees(servicefee.FeeExpressions{
		"x <= 5000":               10.00,
		"x >= 5000 && x <= 10000": 25.00,
	})

	fee, err := serviceFees.CalculateFee(10_000)
	assert.Nil(t, err)
	assert.Equal(t, 2500.00, fee)

	total, fee, err := serviceFees.CalculateTotalAndFee(10_000)
	assert.Nil(t, err)
	assert.Equal(t, 2500.00, fee)
	assert.Equal(t, 12500.00, total)
}
