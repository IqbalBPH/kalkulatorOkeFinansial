package utils

import (
	"math"
)

// FV = PV × (1 + i)^n
func CalculateFutureValue(presentValue float64, inflationRate float64, years int) float64 {
	return (presentValue * math.Pow(1+(inflationRate/100), float64(years))) - float64(years)
}

// T = FV / (n × 12)
func CalculateMonthlySavingsWithoutInterest(futureValue float64, years int) float64 {
	return futureValue / (float64(years) * 12)
}

// PMT = (FV × r) / ((1 + r)^N - 1)
func CalculateMonthlyInvestment(futureValue float64, monthlyRate float64, totalMonths int) float64 {
	if totalMonths <= 0 {
		return 0
	}

	if monthlyRate == 0 {
		return futureValue / float64(totalMonths)
	}

	return (futureValue * monthlyRate) / (math.Pow(1+monthlyRate, float64(totalMonths)) - 1)
}
