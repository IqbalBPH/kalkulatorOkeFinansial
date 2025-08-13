package utils

func CalculateTotalLiquidAssets(liquidAssets []float64) float64 {
	total := 0.0
	for _, asset := range liquidAssets {
		total += asset
	}
	return total
}

func CalculateTotalInvestmentAssets(investmentAssets []float64) float64 {
	total := 0.0
	for _, asset := range investmentAssets {
		total += asset
	}
	return total
}

func CalculateTotalPersonalAssets(personalAssets []float64) float64 {
	total := 0.0
	for _, asset := range personalAssets {
		total += asset
	}
	return total
}

func CalculateTotalAssets(liquidAssets []float64, investmentAssets []float64, personalAssets []float64) float64 {
	totalLiquid := CalculateTotalLiquidAssets(liquidAssets)
	totalInvestment := CalculateTotalInvestmentAssets(investmentAssets)
	totalPersonal := CalculateTotalPersonalAssets(personalAssets)
	return totalLiquid + totalInvestment + totalPersonal
}

func CalculateTotalShortTermLiabilities(shortTermLiabilities []float64) float64 {
	total := 0.0
	for _, liability := range shortTermLiabilities {
		total += liability
	}
	return total
}

func CalculateTotalLongTermLiabilities(longTermLiabilities []float64) float64 {
	total := 0.0
	for _, liability := range longTermLiabilities {
		total += liability
	}
	return total
}

func CalculateTotalLiabilities(liabilities []float64) float64 {
	total := 0.0
	for _, liability := range liabilities {
		total += liability
	}
	return total
}

func CalculateNetWorth(totalAssets float64, totalLiabilities float64) float64 {
	return totalAssets - totalLiabilities
}

// Kas Flow Calculations
func CalculateTotalActiveIncome(activeIncome []float64) float64 {
	total := 0.0
	for _, income := range activeIncome {
		total += income
	}
	return total
}

func CalculateTotalPassiveIncome(passiveIncome []float64) float64 {
	total := 0.0
	for _, income := range passiveIncome {
		total += income
	}
	return total
}

func CalculateTotalIncome(TotalActiveIncome, TotalPassiveIncome float64) float64 {
	return TotalActiveIncome + TotalPassiveIncome
}

func CalculateTotalPosSosial(posSosial []float64) float64 {
	total := 0.0
	for _, pos := range posSosial {
		total += pos
	}
	return total
}

func CalculateTotalPosTabunganInvestasi(posTabunganInvestasi []float64) float64 {
	total := 0.0
	for _, pos := range posTabunganInvestasi {
		total += pos
	}
	return total
}

func CalculateTotalPosBayarKewajibanJangkaPendek(posBayarKewajibanJangkaPendek []float64) float64 {
	total := 0.0
	for _, pos := range posBayarKewajibanJangkaPendek {
		total += pos
	}
	return total
}

func CalculateTotalPosBayarKewajibanJangkaPanjang(posBayarKewajibanJangkaPanjang []float64) float64 {
	total := 0.0
	for _, pos := range posBayarKewajibanJangkaPanjang {
		total += pos
	}
	return total
}

func CalculateTotalPosBayarPremiAsuransi(posBayarPremiAsuransi []float64) float64 {
	total := 0.0
	for _, pos := range posBayarPremiAsuransi {
		total += pos
	}
	return total
}

func CalculateTotalPosBiayaHidupMasaKini(posBiayaHidupMasaKini []float64) float64 {
	total := 0.0
	for _, pos := range posBiayaHidupMasaKini {
		total += pos
	}
	return total
}

func CalculateTotalPengeluaran(annualExpenditures []float64) float64 {
	total := 0.0
	for _, expenditure := range annualExpenditures {
		total += expenditure
	}
	return total
}

func CalculateTotalPendapatan(annualIncome []float64) float64 {
	total := 0.0
	for _, income := range annualIncome {
		total += income
	}
	return total
}

func CalculateArusKas(totalPengeluaran, totalPendapatan float64) float64 {
	return totalPendapatan - totalPengeluaran
}
