package models

type LifeInsurancePlanRequest struct {
	Kpr                  float64 `json:"kpr"`
	Kkb                  float64 `json:"kkb"`
	Kta                  float64 `json:"kta"`
	OnlineLoan           float64 `json:"online_loan"`
	BussinessDebt        float64 `json:"bussiness_debt"`
	OtherDebt            float64 `json:"other_debt"`
	FuneralCost          float64 `json:"funeral_cost"`
	TaxesAndInheritance  float64 `json:"taxes_and_inheritance"`
	OtherCosts           float64 `json:"other_costs"`
	EmergencyFunds       float64 `json:"emergency_funds"`
	HomePurchaseFund     float64 `json:"home_purchase_fund"`
	CarPurchaseFund      float64 `json:"car_purchase_fund"`
	MarriageFund         float64 `json:"marriage_fund"`
	ChildEducationFund   float64 `json:"child_education_fund"`
	WorshipFund          float64 `json:"worship_fund"`
	RetirementFund       float64 `json:"retirement_fund"`
	WakafFund            float64 `json:"wakaf_fund"`
	OtherFinancialFund   float64 `json:"other_financial_fund"`
	IncomeForFamilyNeeds float64 `json:"income_for_family_needs"`
	PeriodOfProtection   int     `json:"period_of_protection"` // dalam tahun
	InvestmentCoverage   float64 `json:"investment_coverage"`  // dalam persentase
	LifeInsuranceOwned   float64 `json:"life_insurance_owned"`
}

type LifeInsurancePlanData struct {
	TotalDebt             int64 `json:"total_debt"`
	TotalFuneralCost      int64 `json:"total_funeral_cost"`
	TotalFamilyFunds      int64 `json:"total_family_funds"`
	NominalSoulCoverage   int64 `json:"nominal_soul_coverage"`
	LifeInsuranceCoverage int64 `json:"life_insurance_coverage"`
	InsuranceMoneyNeeded  int64 `json:"insurance_money_needed"`
}

type LifeInsurancePlanResponse struct {
	Status int                   `json:"status"`
	Data   LifeInsurancePlanData `json:"data"`
}
