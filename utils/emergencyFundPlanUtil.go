package utils

import "kalkulatorOkeFinansial/models"

func CalculateTotalExpenses(ProductiveDebt float64, Insurance float64, CostOfLiving float64, ConsumptiveDebt float64, FutureSavingsAndInvestments float64, LifestyleOrHobbyCosts float64) float64 {
	return ProductiveDebt + Insurance + CostOfLiving + ConsumptiveDebt + FutureSavingsAndInvestments + LifestyleOrHobbyCosts
}

func GetMaritalStatusInt(maritalStatus models.MaritalStatusEnum) int {
	switch maritalStatus {
	case "not_married":
		return 3
	case "married_no_children":
		return 6
	case "married_with_children":
		return 12
	default:
		return 1 // Default case if marital status is unknown
	}
}

func CalculateEmergencyFundNeeded(totalExpenses float64, MaritalStatusInt int) float64 {
	return totalExpenses * float64(MaritalStatusInt)
}

func CalculateLackingFunds(emergencyFundNeeded float64, currentEmergencyFund float64) float64 {
	return emergencyFundNeeded - currentEmergencyFund
}

func CalculateEmergencyFundsSavings(lackingFunds float64, monthsToSave int) float64 {
	if monthsToSave <= 0 {
		return 0
	}
	return lackingFunds / float64(monthsToSave)
}
