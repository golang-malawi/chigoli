package finance

import (
	math "math"
)

// CalculateCompoundInterest calculates the compound interest for a given principal, rate, and years.
// principal: the principal amount
// rate: the rate of interest
// years: the number of years
func CalculateCompoundInterest(principal float64, rate float64, years int) float64 {
    return principal * math.Pow((1 + rate/100), float64(years))
}

// CalculateROI calculates the return on investment for a given initial investment and final value.
// initialInvestment: the initial investment
// finalValue: the final value
func CalculateROI(initialInvestment float64, finalValue float64) float64 {
    return ((finalValue - initialInvestment) / initialInvestment) * 100
}