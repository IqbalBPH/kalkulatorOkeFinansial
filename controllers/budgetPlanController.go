package controllers

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

func CalculateBudgetPlan(c *gin.Context) {
	var request models.BudgetPlanRequest

	// Binding JSON request ke struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Format data tidak valid: " + err.Error(),
		})
		return
	}

	// Validasi input
	if request.ActiveBudget < 0 || request.PassiveBudget < 0 {
		c.JSON(400, gin.H{"error": "Anggaran aktif dan anggaran pasif tidak boleh negatif"})
		return
	}

	// Hitung total anggaran
	totalBudget := utils.CalculateTotalBudget(request.ActiveBudget, request.PassiveBudget)

	// Hitung nilai-nilai finansial lainnya
	data := models.BudgetPlanData{
		TotalBudget:                 int64(totalBudget),
		ProductiveDebt:              int64(utils.CalculateProductiveDebt(totalBudget)),
		Insurance:                   int64(utils.CalculateInsurance(totalBudget)),
		CostOfLiving:                int64(utils.CalculateCostOfLiving(totalBudget)),
		ConsumptiveDebt:             int64(utils.CalculateConsumptiveDebt(totalBudget)),
		FutureSavingsAndInvestments: int64(utils.CalculateFutureSavingsAndInvestments(totalBudget)),
		LifestyleOrHobbyCosts:       int64(utils.CalculateLifestyleOrHobbyCosts(totalBudget)),
	}

	response := models.BudgetPlanResponse{
		Status: 200,
		Data:   data,
	}

	c.JSON(200, response)
}
