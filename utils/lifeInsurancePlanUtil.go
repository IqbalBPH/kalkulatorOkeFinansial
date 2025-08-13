package utils

func CalculateTotalDebt(kpr float64, kkb float64, kta float64, OnlineLoan float64, BussinessDebt float64, OtherDebt float64) float64 {
	return kpr + kkb + kta + OnlineLoan + BussinessDebt + OtherDebt
}

func CalculateTotalFuneralCost(FuneralCost float64, TaxesAndInheritance float64, OtherCosts float64) float64 {
	return FuneralCost + TaxesAndInheritance + OtherCosts
}

func CalculateTotalFamilyFunds(EmergencyFunds float64, HomePurchaseFund float64, CarPurchaseFund float64, MarriageFund float64, ChildEducationFund float64, WorshipFund float64, RetirementFund float64, WakafFund float64, OtherFinancialFund float64) float64 {
	return EmergencyFunds + HomePurchaseFund + CarPurchaseFund + MarriageFund + ChildEducationFund + WorshipFund + RetirementFund + WakafFund + OtherFinancialFund
}

// CalculateNominalSoulCoverage P_keluarga × M_perlindungan
func CalculateNominalSoulCoverage(IncomeForFamilyNeeds float64, PeriodOfProtection int) float64 {
	return IncomeForFamilyNeeds * float64(PeriodOfProtection)
}

// CalculateLifeInsuranceCoverage P_keluarga ÷ (I_bebas risiko ÷ 100)
func CalculateLifeInsuranceCoverage(IncomeForFamilyNeeds float64, InvestmentCoverage float64) float64 {
	if InvestmentCoverage == 0 {
		return 0
	}
	return IncomeForFamilyNeeds / (InvestmentCoverage / 100)
}

// CalculateInsuranceMoneyNeeded uang_pertanggungan_asuransi - up_dimiliki; if (kebutuhan_up < 0) kebutuhan_up = 0
// max(UP_asuransi − UP_dimiliki, 0)
func CalculateInsuranceMoneyNeeded(NominalSoulCoverage float64, LifeInsuranceOwned float64) float64 {
	if NominalSoulCoverage < LifeInsuranceOwned {
		return 0
	}
	return NominalSoulCoverage - LifeInsuranceOwned
}
