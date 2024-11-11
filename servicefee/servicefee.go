package servicefee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/expr-lang/expr"
)

var ErrEmptyExpression = errors.New("expression to be evaluated cannot be empty string")
var ErrExpression = errors.New("the expression to be evaluated is invalid or cannot be evaluated")
var ErrNoApplicableFee = errors.New("no applicable fee was found in the fee table")

type Fees interface {
	CalculateFee(amount float64) (float64, error)

	CalculateTotalAndFee(amount float64) (float64, float64, error)
}

type FeeExpressions map[string]float64

func (f FeeExpressions) Validate() error {
	for k, v := range f {
		if v < 0 {
			return fmt.Errorf("fee cannot be negative")
		}
		switch {
		case strings.Contains(k, ">"):
		case strings.Contains(k, "<"):
		case strings.Contains(k, ">="):
		case strings.Contains(k, "<="):
			continue
		default:
			return fmt.Errorf("expression '%s' does not contain comparison operator", k)
		}

		_, err := evaluateCondition(0, k)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f FeeExpressions) WriteJSON(w io.Writer) error {
	data, err := json.Marshal(f)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

type FixedFees struct {
	feeTable FeeExpressions
}

func evaluateCondition(amount float64, condition string) (bool, error) {
	env := map[string]any{
		"x": amount,
	}

	program, err := expr.Compile(condition, expr.Env(env))
	if err != nil {
		return false, fmt.Errorf("%w: %w", ErrExpression, err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		return false, fmt.Errorf("%w: %w", ErrExpression, err)
	}

	return fmt.Sprint(output) == "true", nil
}

func (f *FixedFees) CalculateFee(amount float64) (float64, error) {
	for expression, fee := range f.feeTable {
		if ok, err := evaluateCondition(amount, expression); err != nil {
			return 0.0, err
		} else if ok {
			return fee, nil
		}
	}

	return 0.0, ErrNoApplicableFee

}

func (f *FixedFees) CalculateTotalAndFee(amount float64) (float64, float64, error) {
	fee, err := f.CalculateFee(amount)
	if err != nil {
		return 0.0, 0.0, err
	}

	return amount + fee, fee, nil
}

type PercentageFees struct {
	feeTable FeeExpressions
}

func NewFixedFee(feeTable FeeExpressions) Fees {
	if feeTable == nil || len(feeTable) < 1 {
		panic(fmt.Errorf("feetable cannot be nil or empty"))
	}
	err := feeTable.Validate()
	if err != nil {
		panic(err)
	}
	return &FixedFees{
		feeTable: feeTable,
	}
}

func NewPercentageFees(feeTable FeeExpressions) Fees {
	if feeTable == nil || len(feeTable) < 1 {
		panic(fmt.Errorf("feetable cannot be nil or empty"))
	}
	err := feeTable.Validate()
	if err != nil {
		panic(err)
	}
	return &PercentageFees{
		feeTable: feeTable,
	}
}

func (f *PercentageFees) CalculateFee(amount float64) (float64, error) {
	for expression, fee := range f.feeTable {
		if ok, err := evaluateCondition(amount, expression); err != nil {
			return 0.0, err
		} else if ok {
			return amount * (fee / 100.00), nil
		}
	}

	return 0.0, ErrNoApplicableFee

}

func (f *PercentageFees) CalculateTotalAndFee(amount float64) (float64, float64, error) {
	fee, err := f.CalculateFee(amount)
	if err != nil {
		return 0.0, 0.0, err
	}

	return amount + fee, fee, nil
}
