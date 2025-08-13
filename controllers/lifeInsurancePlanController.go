package controllers

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

func CalculateLifeInsurancePlan(c *gin.Context) {
	var request models.LifeInsurancePlanRequest

	// Binding JSON request to struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid data format: " + err.Error(),
		})
		return
	}

	TotalDebt := utils.CalculateTotalDebt(request.Kpr, request.Kkb, request.Kta, request.OnlineLoan, request.BussinessDebt, request.OtherDebt)
	TotalFuneralCost := utils.CalculateTotalFuneralCost(request.FuneralCost, request.TaxesAndInheritance, request.OtherCosts)
	TotalFamilyFunds := utils.CalculateTotalFamilyFunds(request.EmergencyFunds, request.HomePurchaseFund, request.CarPurchaseFund, request.MarriageFund, request.ChildEducationFund, request.WorshipFund, request.RetirementFund, request.WakafFund, request.OtherFinancialFund)
	NominalSoulCoverage := utils.CalculateNominalSoulCoverage(request.IncomeForFamilyNeeds, request.PeriodOfProtection)
	LifeInsuranceCoverage := utils.CalculateLifeInsuranceCoverage(request.IncomeForFamilyNeeds, request.InvestmentCoverage)
	InsuranceMoneyNeeded := utils.CalculateInsuranceMoneyNeeded(LifeInsuranceCoverage, request.LifeInsuranceOwned)

	response := models.LifeInsurancePlanResponse{
		Status: 200,
		Data: models.LifeInsurancePlanData{
			TotalDebt:             int64(TotalDebt),
			TotalFuneralCost:      int64(TotalFuneralCost),
			TotalFamilyFunds:      int64(TotalFamilyFunds),
			NominalSoulCoverage:   int64(NominalSoulCoverage),
			LifeInsuranceCoverage: int64(LifeInsuranceCoverage),
			InsuranceMoneyNeeded:  int64(InsuranceMoneyNeeded),
		},
	}

	// Send response
	c.JSON(200, response)
}
