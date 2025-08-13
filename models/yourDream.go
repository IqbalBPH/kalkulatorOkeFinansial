package models

type YourDreamRequest struct {
	PresentValue    float64 `json:"present_value" binding:"required"` // Uang saat ini
	Years           int     `json:"years" binding:"required"`         // Jangka waktu (tahun)
	InflationRate   float64 `json:"inflation_rate"`                   // Asumsi inflasi (desimal)
	InvestmentYield float64 `json:"investment_yield"`                 // Asumsi hasil investasi (desimal)
}

// Data hasil perhitungan
type YourDreamData struct {
	FutureValue             int64 `json:"future_value"`              // Nilai masa depan
	MonthlySavings          int64 `json:"monthly_savings"`           // Tabungan bulanan
	MonthlyInvestmentAmount int64 `json:"monthly_investment_amount"` // Investasi bulanan
}

type YourDreamResponse struct {
	Status int           `json:"status"` // Status dari permintaan
	Data   YourDreamData `json:"data"`   // Data hasil perhitungan
}
