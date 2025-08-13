package utils

func CalculateTotalBudget(activeBudget, passiveBudget float64) float64 {
	return activeBudget + passiveBudget
}

func CalculateProductiveDebt(totalBudget float64) float64 {
	return 0.20 * totalBudget
}

func CalculateInsurance(totalBudget float64) float64 {
	return 0.10 * totalBudget
}

func CalculateCostOfLiving(totalBudget float64) float64 {
	return 0.40 * totalBudget
}

func CalculateConsumptiveDebt(totalBudget float64) float64 {
	return 0.15 * totalBudget
}

func CalculateFutureSavingsAndInvestments(totalBudget float64) float64 {
	return 0.10 * totalBudget
}
func CalculateLifestyleOrHobbyCosts(totalBudget float64) float64 {
	return 0.05 * totalBudget
}
