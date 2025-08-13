package models

type BudgetPlanRequest struct {
	ActiveBudget  float64 `json:"active_budget"`
	PassiveBudget float64 `json:"passive_budget"`
}

// Data hasil perhitungan
type BudgetPlanData struct {
	TotalBudget                 int64 `json:"total_budget"`
	ProductiveDebt              int64 `json:"productive_debt"`
	Insurance                   int64 `json:"insurance"`
	CostOfLiving                int64 `json:"cost_of_living"`
	ConsumptiveDebt             int64 `json:"consumptive_debt"`
	FutureSavingsAndInvestments int64 `json:"future_savings_and_investments"`
	LifestyleOrHobbyCosts       int64 `json:"lifestyle_or_hobby_costs"`
}

type BudgetPlanResponse struct {
	Status int            `json:"status"` // Status dari permintaan
	Data   BudgetPlanData `json:"data"`   // Data hasil perhitungan
}
