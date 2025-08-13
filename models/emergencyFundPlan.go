package models

type MaritalStatusEnum string

const (
	NotMarried       MaritalStatusEnum = "not_married"
	MarriedNoChild   MaritalStatusEnum = "married_no_children"
	MarriedWithChild MaritalStatusEnum = "married_with_children"
)

type EmergencyFundPlanRequest struct {
	ProductiveDebt              float64           `json:"productive_debt"`
	ConsumptiveDebt             float64           `json:"consumptive_debt"`
	PremiInsurance              float64           `json:"premi_insurance"`
	FutureSavingsAndInvestments float64           `json:"future_savings_and_investments"`
	CostOfLiving                float64           `json:"cost_of_living"`
	LifestyleOrHobbyCosts       float64           `json:"lifestyle_or_hobby_costs"`
	MaritalStatus               MaritalStatusEnum `json:"marital_status"`
	EmergencyFundsowned         float64           `json:"emergency_funds_owned"`
	TimePeriod                  int               `json:"time_period"`
}

// Data hasil perhitungan
type EmergencyFundPlanData struct {
	TotalExpenses         int64 `json:"total_expenses"`
	EmergencyFundNeeded   int64 `json:"emergency_fund_needed"`
	LackingFunds          int64 `json:"lacking_funds"`
	EmergencyFundsSavings int64 `json:"emergency_funds_savings"`
}

type EmergencyFundPlanResponse struct {
	Status int                   `json:"status"` // Status dari permintaan
	Data   EmergencyFundPlanData `json:"data"`   // Data hasil perhitungan
}
