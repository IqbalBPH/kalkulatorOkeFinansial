package models

type HomePurchasePlanRequest struct {
	HomePrice              float64 `json:"home_price"`
	PurchasePeriod         int     `json:"purchase_period"`      // dalam tahun
	InflationRate          float64 `json:"inflation_rate"`       // dalam persentase
	PercentDownPayment     float64 `json:"percent_down_payment"` // dalam persentase
	NotaryCost             float64 `json:"notary_cost"`
	KprAdminCost           float64 `json:"kpr_admin_cost"`
	Ppn                    float64 `json:"ppn"`
	KprInsuranceCost       float64 `json:"kpr_insurance_cost"`
	AnnualInvestmentTarget float64 `json:"annual_investment_target"` // dalam persentase
	KprInterestRate        float64 `json:"kpr_interest_rate"`        // dalam persentase
	LoanTenor              int     `json:"loan_tenor"`               // dalam tahun
}

type HomePurchasePlanData struct {
	TotalHomePrice      int64 `json:"total_home_price"`      // Total harga rumah
	DownPayment         int64 `json:"down_payment"`          // Uang muka
	TotalDownPayment    int64 `json:"total_down_payment"`    // Total uang muka
	SavingPerMonth      int64 `json:"saving_per_month"`      // Tabungan per bulan
	MonthlyInvestment   int64 `json:"monthly_investment"`    // Investasi bulanan
	KprNominal          int64 `json:"kpr_nominal"`           // Nominal KPR
	InstallmentPerMonth int64 `json:"installment_per_month"` // Cicilan per bulan
	TotalInterestCost   int64 `json:"total_interest_cost"`   // Total biaya bunga
	TotalCreditInterest int64 `json:"total_credit_interest"` // Total bunga kredit
}

type HomePurchasePlanResponse struct {
	Status int                  `json:"status"` // Status dari permintaan
	Data   HomePurchasePlanData `json:"data"`   // Data hasil perhitungan
}
