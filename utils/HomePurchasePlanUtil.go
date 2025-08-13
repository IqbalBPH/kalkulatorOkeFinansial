package utils

import "math"

func FV(rate float64, nper int, pmt float64, pv float64, typ int) float64 {
	if nper <= 0 {
		return -pv
	}
	if rate != 0 {
		pow := math.Pow(1+rate, float64(nper))
		return (pmt * (1 + rate*float64(typ)) * (1 - pow) / rate) - pv*pow
	}
	return -(pv + pmt*float64(nper))
}

func PMT(rate float64, nper int, pv float64, fv float64, typ int) float64 {
	if nper == 0 {
		return 0
	}
	if rate != 0 {
		q := math.Pow(1+rate, float64(nper))
		den := (-1 + q) * (1 + rate*float64(typ))
		if den == 0 {
			return 0
		}
		return -(rate * (fv + q*pv)) / den
	}
	return -(fv + pv) / float64(nper)
}

func TotalHomePrice(homePrice float64, purchasePeriodYears int, inflationPct float64) float64 {
	r := inflationPct / 100.0
	return math.Round(FV(r, purchasePeriodYears, 1, -homePrice, 0))
}

func CalculateDownPayment(totalHomePriceRounded float64, percentDownPayment float64) float64 {
	if totalHomePriceRounded <= 0 {
		return 0
	}
	if percentDownPayment < 0 {
		percentDownPayment = 0
	}
	return math.Round(totalHomePriceRounded * (percentDownPayment / 100.0))
}

func CalculateTotalDownPayment(downPayment, notaryCost, kprAdminCost, ppn, kprInsuranceCost float64) float64 {
	if downPayment < 0 {
		downPayment = 0
	}
	if notaryCost < 0 {
		notaryCost = 0
	}
	if kprAdminCost < 0 {
		kprAdminCost = 0
	}
	if ppn < 0 {
		ppn = 0
	}
	if kprInsuranceCost < 0 {
		kprInsuranceCost = 0
	}
	return downPayment + notaryCost + kprAdminCost + ppn + kprInsuranceCost
}

func CalculateSavingPerMonth(totalDownPayment float64, purchasePeriodYears int) float64 {
	if totalDownPayment <= 0 || purchasePeriodYears <= 0 {
		return 0
	}
	return math.Round(totalDownPayment / float64(purchasePeriodYears*12))
}

func CalculateMonthlyHomeInvestment(annualInvestmentTargetPct float64, purchasePeriodYears int, totalDownPayment float64) float64 {
	if purchasePeriodYears <= 0 || totalDownPayment <= 0 {
		return 0
	}
	r := annualInvestmentTargetPct / 1200.0
	n := purchasePeriodYears * 12
	// PMT(r, n, 0, -totalDownPayment, 0) lalu dibulatkan
	return math.Round(PMT(r, n, 0, -totalDownPayment, 0))
}

func CalculateKprNominal(totalHomePriceRounded float64, downPaymentRounded float64) float64 {
	// nominal_kpr = harga_terbaru - uang_muka (tanpa biaya lain)
	if totalHomePriceRounded <= 0 {
		return 0
	}
	if downPaymentRounded < 0 {
		downPaymentRounded = 0
	}
	return totalHomePriceRounded - downPaymentRounded
}

func CalculateInstallmentPerMonth(kprNominal float64, kprInterestRateAnnualPct float64, loanTenorYears int) float64 {
	if kprNominal <= 0 || loanTenorYears <= 0 {
		return 0
	}
	r := kprInterestRateAnnualPct / 1200.0
	n := loanTenorYears * 12
	// PMT(r, n, -kprNominal, 0, 0) lalu dibulatkan untuk display
	return math.Round(PMT(r, n, -kprNominal, 0, 0))
}

func CalculateTotals(installmentPerMonth float64, loanTenorYears int, kprNominal float64) (totalPaidRounded float64, totalInterestRounded float64) {
	if installmentPerMonth <= 0 || loanTenorYears <= 0 || kprNominal <= 0 {
		return 0, 0
	}
	n := float64(loanTenorYears * 12)
	totalPaid := installmentPerMonth * n // nilai cicilan (float) * n
	totalInterest := totalPaid - kprNominal
	return math.Round(totalPaid), math.Round(totalInterest)
}
