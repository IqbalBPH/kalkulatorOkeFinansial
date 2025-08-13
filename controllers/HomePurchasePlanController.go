package controllers

import (
	"math"

	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

func CalculateHomePurchasePlan(c *gin.Context) {
	var request models.HomePurchasePlanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data format: " + err.Error()})
		return
	}

	totalHomePrice := utils.TotalHomePrice(request.HomePrice, request.PurchasePeriod, request.InflationRate)
	downPayment := utils.CalculateDownPayment(totalHomePrice, request.PercentDownPayment)
	totalDownPayment := utils.CalculateTotalDownPayment(
		downPayment,
		request.NotaryCost,
		request.KprAdminCost,
		request.Ppn,
		request.KprInsuranceCost,
	)
	savingPerMonth := utils.CalculateSavingPerMonth(totalDownPayment, request.PurchasePeriod)
	monthlyHomeInvestment := utils.CalculateMonthlyHomeInvestment(
		request.AnnualInvestmentTarget, // % per tahun
		request.PurchasePeriod,
		totalDownPayment,
	)
	kprNominal := utils.CalculateKprNominal(totalHomePrice, downPayment)
	installmentPerMonth := utils.CalculateInstallmentPerMonth(kprNominal, request.KprInterestRate, request.LoanTenor)
	totalPaid, totalInterest := utils.CalculateTotals(installmentPerMonth, request.LoanTenor, kprNominal)

	// Helper pembulatan ke integer
	toInt := func(v float64) int64 { return int64(math.Round(v)) }

	response := models.HomePurchasePlanResponse{
		Status: 200,
		Data: models.HomePurchasePlanData{
			TotalHomePrice:      toInt(totalHomePrice),
			DownPayment:         toInt(downPayment),
			TotalDownPayment:    toInt(totalDownPayment),
			SavingPerMonth:      toInt(savingPerMonth),
			MonthlyInvestment:   toInt(monthlyHomeInvestment),
			KprNominal:          toInt(kprNominal),
			InstallmentPerMonth: toInt(installmentPerMonth),
			TotalInterestCost:   toInt(totalPaid),
			TotalCreditInterest: toInt(totalInterest),
		},
	}

	c.JSON(200, response)
}
