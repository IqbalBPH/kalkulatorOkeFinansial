package controllers

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

func CalculateEmergencyFund(c *gin.Context) {
	var request models.EmergencyFundPlanRequest

	// Binding JSON request ke struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Format data not valid: " + err.Error(),
		})
		return
	}

	totalExpenses := utils.CalculateTotalExpenses(request.ProductiveDebt, request.ConsumptiveDebt, request.PremiInsurance, request.FutureSavingsAndInvestments, request.CostOfLiving, request.LifestyleOrHobbyCosts)
	// get marital status as integer
	maritalStatusInt := utils.GetMaritalStatusInt(request.MaritalStatus)
	emergencyFundNeeded := utils.CalculateEmergencyFundNeeded(totalExpenses, maritalStatusInt)
	lackingFunds := utils.CalculateLackingFunds(emergencyFundNeeded, request.EmergencyFundsowned)
	emergencyFundsSavings := utils.CalculateEmergencyFundsSavings(lackingFunds, request.TimePeriod)

	// Prepare response data
	response := models.EmergencyFundPlanResponse{
		Status: 200,
		Data: models.EmergencyFundPlanData{
			TotalExpenses:         int64(totalExpenses),
			EmergencyFundNeeded:   int64(emergencyFundNeeded),
			LackingFunds:          int64(lackingFunds),
			EmergencyFundsSavings: int64(emergencyFundsSavings),
		},
	}

	// Send response
	c.JSON(200, response)

}
