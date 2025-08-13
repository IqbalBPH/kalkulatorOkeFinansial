package models

type RetireFundPlanRequest struct {
	CurrentAge               int     `json:"current_age"`                 // Usia saat ini
	RetirementAge            int     `json:"retirement_age"`              // Usia pens
	AgeRetireFundNeeded      int     `json:"age_retire_fund_needed"`      // Usia saat dana pensiun dibutuhkan
	CurrentMonthlyIncome     float64 `json:"current_monthly_income"`      // Penghasilan bulanan saat ini
	PersentDesiredMonthlyFee float64 `json:"persent_desired_monthly_fee"` // Persentase penghasilan yang diinginkan saat pensiun
	InflationRate            float64 `json:"inflation_rate"`              // Tingkat infl (persentase)
	InvestmentReturnRate     float64 `json:"investment_return_rate"`      // Tingkat pengembalian investasi (persentase)
}

type RetireFundPlanData struct {
	RetireYearDiff          int64 `json:"retire_year_diff"`          // Selisih waktu tahun pensiun yang dibutuhkan
	LifeExpectancy          int64 `json:"life_expectancy"`           // Harapan hidup
	MonthlyRetirementIncome int64 `json:"monthly_retirement_income"` // Penghasilan bulanan saat pensiun
	AnnualRetirementIncome  int64 `json:"annual_retirement_income"`  // Penghasilan tahunan saat pensiun
	MonthlyFutureValue      int64 `json:"monthly_future_value"`      // Nilai masa depan bulanan
	AnnualFutureValue       int64 `json:"annual_future_value"`       // Nilai masa depan tahunan
	AmountRetirementFund    int64 `json:"amount_retirement_fund"`    // Jumlah dana pensiun yang dibutuhkan
	MoneySaving             int64 `json:"money_saving"`              // Tabungan yang harus disisihkan
	MonthlyInvest           int64 `json:"monthly_investment_amount"` // Jumlah investasi bulanan
}

type RetireFundPlanResponse struct {
	Status int                `json:"status"` // Status dari permintaan
	Data   RetireFundPlanData `json:"data"`   // Data hasil perhitungan
}
