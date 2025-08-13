package controllers

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
	"math"
)

func CalculateRetireFundPlan(c *gin.Context) {
	var request models.RetireFundPlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	// Calculate the retirement fund plan
	yearDiff := utils.CalculateRetireYearDiff(request.CurrentAge, request.RetirementAge)
	LifeExpectancy := utils.CalculateLifeExpectancy(request.RetirementAge, request.AgeRetireFundNeeded)
	monthlyRetirementIncome := utils.CalculateMonthlyRetirementIncome(request.CurrentMonthlyIncome, request.PersentDesiredMonthlyFee)
	annualRetirementIncome := utils.CalculateAnnualRetirementIncome(monthlyRetirementIncome)
	MonthlyFutureValue := utils.CalculateMonthlyFutureValue(
		request.InflationRate,
		yearDiff,
		float64(monthlyRetirementIncome),
	)
	AnnualFutureValue := utils.CalculateAnnualFutureValue(
		request.InflationRate,
		yearDiff,
		float64(annualRetirementIncome),
	)

	AmountRetirementFund := utils.CalculateRetirementFundNeeded(AnnualFutureValue, LifeExpectancy)
	MoneySaving := utils.CalculateMonthlySavingSimple(AmountRetirementFund, yearDiff)
	MonthlyInvest := utils.CalculateMonthlyRetireInvestment(request.InvestmentReturnRate, yearDiff, AmountRetirementFund)

	toInt := func(v float64) int64 { return int64(math.Round(v)) }

	resp := models.RetireFundPlanResponse{
		Status: 200,
		Data: models.RetireFundPlanData{
			RetireYearDiff:          yearDiff,
			LifeExpectancy:          LifeExpectancy,
			MonthlyRetirementIncome: monthlyRetirementIncome,
			AnnualRetirementIncome:  annualRetirementIncome,
			MonthlyFutureValue:      toInt(MonthlyFutureValue),
			AnnualFutureValue:       toInt(AnnualFutureValue),
			AmountRetirementFund:    toInt(AmountRetirementFund),
			MoneySaving:             toInt(MoneySaving),
			MonthlyInvest:           toInt(MonthlyInvest),
		},
	}

	c.JSON(200, resp)

	// json

}
