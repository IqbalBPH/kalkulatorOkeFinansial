package utils

func CalculateRetireYearDiff(currentAge, retirementAge int) int64 {
	if currentAge >= retirementAge {
		return 0
	}
	return int64(retirementAge - currentAge)
}

func CalculateLifeExpectancy(RetirementAge, AgeRetireFundNeeded int) int64 {
	return int64(AgeRetireFundNeeded - RetirementAge)
}

func CalculateMonthlyRetirementIncome(currentMonthlyIncome, persentDesiredMonthlyFee float64) int64 {
	if persentDesiredMonthlyFee < 0 {
		return 0
	}
	return int64(currentMonthlyIncome * (persentDesiredMonthlyFee / 100))
}

func CalculateAnnualRetirementIncome(monthlyRetirementIncome int64) int64 {
	if monthlyRetirementIncome < 0 {
		return 0
	}
	return monthlyRetirementIncome * 12
}

// FV_bulanan = FV(inflasi/100, kapan_pensiun, 1, -biaya_bulanan, 0)
func CalculateMonthlyFutureValue(inflationPct float64, yearsToRetire int64, monthlyExpense float64) float64 {
	if yearsToRetire <= 0 || monthlyExpense <= 0 {
		return 0
	}
	r := inflationPct / 100.0
	return FV(r, int(yearsToRetire), 1, -monthlyExpense, 0)
}

// FV_tahunan = FV(inflasi/100, kapan_pensiun, 1, -biaya_tahunan, 0)
func CalculateAnnualFutureValue(inflationPct float64, yearsToRetire int64, annualExpense float64) float64 {
	if yearsToRetire <= 0 || annualExpense <= 0 {
		return 0
	}
	r := inflationPct / 100.0
	return FV(r, int(yearsToRetire), 1, -annualExpense, 0)
}

// Dana_pensiun = FV_pengeluaran_tahunan × Estimasi_harapan_hidup
func CalculateRetirementFundNeeded(annualFutureValue float64, lifeExpectancyYears int64) float64 {
	if annualFutureValue <= 0 || lifeExpectancyYears <= 0 {
		return 0
	}
	return annualFutureValue * float64(lifeExpectancyYears)
}

// Tabungan_perbulan = Dana_pensiun / (kapan_pensiun × 12)
func CalculateMonthlySavingSimple(retirementFund float64, yearsToRetire int64) float64 {
	if retirementFund <= 0 || yearsToRetire <= 0 {
		return 0
	}
	return retirementFund / float64(yearsToRetire*12)
}

// Investasi_perbulan = PMT(investasi/1200, kapan_pensiun × 12, 0, -Dana_pensiun, 0)
func CalculateMonthlyRetireInvestment(investmentReturnPct float64, yearsToRetire int64, retirementFund float64) float64 {
	if retirementFund <= 0 || yearsToRetire <= 0 {
		return 0
	}
	r := investmentReturnPct / 1200.0
	n := int(yearsToRetire * 12)
	return PMT(r, n, 0, -retirementFund, 0)
}
