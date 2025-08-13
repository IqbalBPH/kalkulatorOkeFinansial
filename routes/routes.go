package routes

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/controllers"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "OkeFinansial Calculator API aktif dan berjalan!",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/v1/yourDream", controllers.CalculateYourDream)
		api.POST("/v1/budgetPlan", controllers.CalculateBudgetPlan)
		api.POST("/v1/emergencyFundPlan", controllers.CalculateEmergencyFund)
		api.POST("/v1/lifeInsurancePlan", controllers.CalculateLifeInsurancePlan)
		api.POST("/v1/homePurchasePlan", controllers.CalculateHomePurchasePlan)
		api.POST("/v1/retireFundPlan", controllers.CalculateRetireFundPlan)
		// Financial Checkup
		//api.POST("/v1/financialCheckup", controllers.CalculateFinancialCheckup)
		api.POST("/v1/assetAndLiability", controllers.CalculateAssetAndLiability)
		api.POST("/v1/kasFlow", controllers.CalculateKasFlow)
	}
}
