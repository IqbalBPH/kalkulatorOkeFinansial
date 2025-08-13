package controllers

import (
	"github.com/gin-gonic/gin"
	"kalkulatorOkeFinansial/models"
	"kalkulatorOkeFinansial/utils"
)

func CalculateAssetAndLiability(c *gin.Context) {
	var request models.AssetAndLiabilityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	LiquidAssets := []float64{request.KasDiTangan, request.Tabungan, request.Deposito, request.ReksadanaPasarUang}
	InvestmentAssets := []float64{request.EmasLogamBerharga, request.ReksadanaPendapatanTetap, request.ReksadanaCampuran, request.ReksadanaSaham, request.Obligasi, request.Saham, request.NilaiTunaiPolisAsuransi, request.DanaPensiunLembagaKeuangan, request.BpjsTenagaKerja, request.BarangKoleksi, request.PropertiRumahSewa, request.Tanah, request.NilaiBersihUsaha, request.AsetLainnya}
	PersonalAssets := []float64{request.RumahDitempati, request.Perhiasan, request.Mobil, request.Motor, request.LaptopHandphone, request.AsetPersonalLainnya}
	ShortTermLiabilities := []float64{request.KewajibanKartuKredit, request.PinjamanPribadi, request.PinjamanMobil, request.PinjamanKta, request.PendekPinjamanLainnya}
	LongTermLiabilities := []float64{request.PinjamanRumah, request.PinjamanApartemen, request.PinjamanLunakPerusahaan, request.PanjangPinjamanLainnya}

	totalLiquid := utils.CalculateTotalLiquidAssets(LiquidAssets)
	totalInvestment := utils.CalculateTotalInvestmentAssets(InvestmentAssets)
	totalPersonal := utils.CalculateTotalPersonalAssets(PersonalAssets)
	totalShortTerm := utils.CalculateTotalShortTermLiabilities(ShortTermLiabilities)
	totalLongTerm := utils.CalculateTotalLongTermLiabilities(LongTermLiabilities)
	liabilities := append(ShortTermLiabilities, LongTermLiabilities...)

	totalAssets := utils.CalculateTotalAssets(
		LiquidAssets,
		InvestmentAssets,
		PersonalAssets,
	)
	totalLiabilities := utils.CalculateTotalLiabilities(
		liabilities,
	)

	netWorth := utils.CalculateNetWorth(
		totalAssets,
		totalLiabilities,
	)

	resp := models.AssetAndLiabilityResponse{
		Status: 200,
		Data: models.AssetAndLiabilityData{
			TotalLiquidAssets:         int64(totalLiquid),
			TotalInvestmentAssets:     int64(totalInvestment),
			TotalPersonalAssets:       int64(totalPersonal),
			TotalAssets:               int64(totalAssets),
			TotalShortTermLiabilities: int64(totalShortTerm),
			TotalLongTermLiabilities:  int64(totalLongTerm),
			TotalLiabilities:          int64(totalLiabilities),
			TotalNetWorth:             int64(netWorth),
		},
	}

	c.JSON(200, resp)
}

func CalculateKasFlow(c *gin.Context) {
	var request models.KasFlowRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	TotalActiveIncomeArray := []float64{
		request.Gaji,
		request.Tunjangan,
		request.InsentifBonus,
		request.PendapatanBisnis,
		request.PendapatanLainnya,
	}
	TotalPassiveIncomeArray := []float64{
		request.BungaDeposito,
		request.KuponObligasi,
		request.DividenSaham,
		request.PendapatanSewa,
		request.PendapatanPasifLainnya,
	}
	totalActiveIncome := utils.CalculateTotalActiveIncome(TotalActiveIncomeArray)
	totalPassiveIncome := utils.CalculateTotalPassiveIncome(TotalPassiveIncomeArray)
	totalIncome := utils.CalculateTotalIncome(totalActiveIncome, totalPassiveIncome)

	resp := models.KasFlowResponse{
		Status: 200,
		Data: models.KasFlowData{
			TotalActiveIncome:  int64(totalActiveIncome),
			TotalPassiveIncome: int64(totalPassiveIncome),
			TotalIncome:        int64(totalIncome),
		},
	}

	c.JSON(200, resp)
}

//func CalculateFinancialCheckup(c *gin.Context) {
//	var request models.FinancialCheckUpRequest
//
//	if err := c.ShouldBindJSON(&request); err != nil {
//		c.JSON(400, gin.H{
//			"status":  400,
//			"message": "Invalid request data",
//			"error":   err.Error(),
//		})
//		return
//	}
//
//	// Calculate Asset and Liability
//	assetAndLiabilityResp := CalculateAssetAndLiabilityData(request.AssetAndLiabilityRequest)
//
//	// Calculate Kas Flow
//	kasFlowResp := CalculateKasFlowData(request.KasFlowRequest)
//
//	// Calculate Annual Expend
//	annualExpendResp := CalculateAnnualExpendData(request.AnnualExpendRequest)
//
//	resp := models.FinancialCheckUpResponse{
//		Status: 200,
//		Data: models.FinancialCheckUpData{
//			AssetAndLiabilityData: assetAndLiabilityResp,
//			KasFlowData:           kasFlowResp,
//			AnnualExpendData:      annualExpendResp,
//		},
//	}
//
//	c.JSON(200, resp)
//}
