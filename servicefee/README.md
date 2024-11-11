# servicefee

A simple library for calculating a fee or amount - supports fixed fees and percentage based fees.

Useful for
  - calculating transaction fees for online transactions
  - calculating administrative fees
  

## Fixed Fees Table

Fixed fees allows for fees to be calculated based the amount, the fees themselves are fixed.

```go
package main

import "github.com/golang-malawi/chigoli/servicefee"

func main() {
	serviceFees := servicefee.NewFixedFees(servicefee.FeeExpressions{
		"x <= 5000":               200.00,
		"x >= 5000 && x <= 10000": 500.00,
	})

	fee, _ := serviceFees.CalculateFee(10_000)
    fmt.Println(fee)

	total, fee, _ := serviceFees.CalculateTotalAndFee(10_000)
    fmt.Println("total", total, "fee", fee)
}
```

## Percentage based Fees Table

Percentage-based fees allows for fees to be calculated based on percentage of the amount.
This is typical for additional administrative fees.

```go
package main

import "github.com/golang-malawi/chigoli/servicefee"

func main() {
	serviceFees := servicefee.NewPercentageFees(servicefee.FeeExpressions{
		"x <= 5000":               10.00,
		"x >= 5000 && x <= 10000": 25.00,
	})

	fee, _ := serviceFees.CalculateFee(10_000)
    fmt.Println(fee)

	total, fee, _ := serviceFees.CalculateTotalAndFee(10_000)
    fmt.Println("total", total, "fee", fee)
}
```