package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

// Calculate menangani permintaan untuk menghitung nilai-nilai finansial
func CalculateYourDream(c *gin.Context) {
	var request models.YourDreamRequest

	// Binding JSON request ke struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format data tidak valid: " + err.Error(),
		})
		return
	}

	// Validasi input
	if request.PresentValue <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nilai uang saat ini harus lebih dari 0"})
		return
	}

	if request.Years <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Jangka waktu harus lebih dari 0"})
		return
	}

	if request.InflationRate < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tingkat inflasi tidak boleh negatif"})
		return
	}

	if request.InvestmentYield < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hasil investasi tidak boleh negatif"})
		return
	}

	// Hitung Future Value berdasarkan inflasi
	futureValue := utils.CalculateFutureValue(
		request.PresentValue,
		request.InflationRate,
		request.Years,
	)

	// Hitung tabungan bulanan tanpa bunga
	monthlySavings := utils.CalculateMonthlySavingsWithoutInterest(
		futureValue,
		request.Years,
	)

	// Hitung investasi bulanan dengan bunga
	monthlyRate := (request.InvestmentYield / 100.0) / 12.0
	totalMonths := request.Years * 12
	monthlyInvestment := utils.CalculateMonthlyInvestment(
		futureValue,
		monthlyRate,
		totalMonths,
	)

	response := models.YourDreamResponse{
		Status: 200,
		Data: models.YourDreamData{
			FutureValue:             int64(futureValue),
			MonthlySavings:          int64(monthlySavings),
			MonthlyInvestmentAmount: int64(monthlyInvestment),
		},
	}

	// Kirim respons
	c.JSON(http.StatusOK, response)
}
