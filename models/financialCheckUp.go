package models

type AssetAndLiabilityRequest struct {
	// ASET LIKUID
	KasDiTangan                float64 `json:"kas_di_tangan"`
	Tabungan                   float64 `json:"tabungan"`
	Deposito                   float64 `json:"deposito"`
	ReksadanaPasarUang         float64 `json:"reksadana_pasar_uang"`
	EmasLogamBerharga          float64 `json:"emas_logam_berharga"` // ASET INVESTASI
	ReksadanaPendapatanTetap   float64 `json:"reksadana_pendapatan_tetap"`
	ReksadanaCampuran          float64 `json:"reksadana_campuran"`
	ReksadanaSaham             float64 `json:"reksadana_saham"`
	Obligasi                   float64 `json:"obligasi"`
	Saham                      float64 `json:"saham"`
	NilaiTunaiPolisAsuransi    float64 `json:"nilai_tunai_polis_asuransi"`
	DanaPensiunLembagaKeuangan float64 `json:"dana_pensiun_lembaga_keuangan"`
	BpjsTenagaKerja            float64 `json:"bpjs_tenaga_kerja"`
	BarangKoleksi              float64 `json:"barang_koleksi"`
	PropertiRumahSewa          float64 `json:"properti_rumah_sewa"`
	Tanah                      float64 `json:"tanah"`
	NilaiBersihUsaha           float64 `json:"nilai_bersih_usaha"`
	AsetLainnya                float64 `json:"aset_lainnya"`
	RumahDitempati             float64 `json:"rumah_ditempati"` // ASET PERSONAL
	Perhiasan                  float64 `json:"perhiasan"`
	Mobil                      float64 `json:"mobil"`
	Motor                      float64 `json:"motor"`
	LaptopHandphone            float64 `json:"laptop_handphone"`
	AsetPersonalLainnya        float64 `json:"aset_personal_lainnya"`

	// KEWAJIBAN JANGKA PENDEK
	KewajibanKartuKredit    float64 `json:"kewajiban_kartu_kredit"`
	PinjamanPribadi         float64 `json:"pinjaman_pribadi"`
	PinjamanMobil           float64 `json:"pinjaman_mobil"`
	PinjamanKta             float64 `json:"pinjaman_kta"`
	PendekPinjamanLainnya   float64 `json:"pendek_pinjaman_lainnya"`
	PinjamanRumah           float64 `json:"pinjaman_rumah"` // KEWAJIBAN JANGKA PANJANG
	PinjamanApartemen       float64 `json:"pinjaman_apartemen"`
	PinjamanLunakPerusahaan float64 `json:"pinjaman_lunak_perusahaan"`
	PanjangPinjamanLainnya  float64 `json:"panjang_pinjaman_lainnya"`
}

type KasFlowRequest struct {
	// PENDAPATAN AKTIF
	Gaji              float64 `json:"gaji"`
	Tunjangan         float64 `json:"tunjangan"`
	InsentifBonus     float64 `json:"insentif_bonus"`
	PendapatanBisnis  float64 `json:"pendapatan_bisnis"`
	PendapatanLainnya float64 `json:"pendapatan_lainnya"`

	// PENDAPATAN PASIF
	BungaDeposito          float64 `json:"bunda_deposito"`
	KuponObligasi          float64 `json:"kupon_obligasi"`
	DividenSaham           float64 `json:"dividen_saham"`
	PendapatanSewa         float64 `json:"pendapatan_sewa"`
	PendapatanPasifLainnya float64 `json:"pendapatan_pasif_lainnya"`
}

type AnnualExpendRequest struct {
	// POS SOSIAL
	Zakat            float64 `json:"zakat"`
	InfaqSedekah     float64 `json:"infaq_sedekah"`
	Perpuluhan       float64 `json:"perpuluhan"`
	Sumbangan        float64 `json:"sumbangan"`
	PosSosialLainnya float64 `json:"pos_sosial_lainnya"`

	// POS TABUNGAN & INVESTASI
	TabunganDanaDarurat        float64 `json:"tabungan_dana_darurat"`
	TabunganPembelianRumah     float64 `json:"tabungan_pembelian_rumah"`
	TabunganPembelianKendaraan float64 `json:"tabungan_pembelian_kendaraan"`
	TabunganPernikahan         float64 `json:"tabungan_pernikahan"`
	TabunganPendidikanAnak     float64 `json:"tabungan_pendidikan_anak"`
	TabunganDanaIbadah         float64 `json:"tabungan_dana_ibadah"`
	TabunganDanaPensiun        float64 `json:"tabungan_dana_pensiun"`
	TabunganLainnya            float64 `json:"tabungan_lainnya"`

	// POS BAYAR KEWAJIBAN JANGKA PENDEK
	PosBayarKartuKredit      float64 `json:"pos_bayar_kartu_kredit"`
	PosPinjamanPribadi       float64 `json:"pos_pinjaman_pribadi"`
	PosPinjamanMobil         float64 `json:"pos_pinjaman_mobil"`
	PosPinjamanKta           float64 `json:"pos_pinjaman_kta"`
	PendekPosPinjamanLainnya float64 `json:"pendek_pos_pinjaman_lainnya"`

	// POS BAYAR KEWAJIBAN JANGKA PANJANG
	PosPinjamanRumah           float64 `json:"pos_pinjaman_rumah"`
	PosPinjamanApartemen       float64 `json:"pos_pinjaman_apartemen"`
	PosPinjamanLunakPerusahaan float64 `json:"pos_pinjaman_lunak_perusahaan"`
	PanjangPosPinjamanLainnya  float64 `json:"panjang_pos_pinjaman_lainnya"`

	// POS BAYAR PREMI ASURANSI
	BpjsKesehatan          float64 `json:"bpjs_kesehatan"`
	AsuransiJiwa           float64 `json:"asuransi_jiwa"`
	AsuransiKesehatan      float64 `json:"asuransi_kesehatan"`
	AsuransiPenyakitKrisis float64 `json:"asuransi_penyakit_krisis"`
	AsuransiRumah          float64 `json:"asuransi_rumah"`
	AsuransiKendaraan      float64 `json:"asuransi_kendaraan"`

	// POS BIAYA HIDUP MASA KINI
	BelanjaRumahTangga float64 `json:"belanja_rumah_tangga"`
	Utilitas           float64 `json:"utilitas"`
	Pakaian            float64 `json:"pakaian"`
	MakananMinuman     float64 `json:"makanan_minuman"`
	Transportasi       float64 `json:"transportasi"`
	Anak               float64 `json:"anak"`
	Kesehatan          float64 `json:"kesehatan"`
	Pendidikan         float64 `json:"pendidikan"`
	Kendaraan          float64 `json:"kendaraan"`
	Hobi               float64 `json:"hobi"`
	GajiArt            float64 `json:"gaji_art"`
	Hiburan            float64 `json:"hiburan"`
	GayaHidup          float64 `json:"gaya_hidup"`
	PengeluaranLainnya float64 `json:"pengeluaran_lainnya"`
}

type FinancialCheckUpRequest struct {
	AssetAndLiabilityRequest
	KasFlowRequest
	AnnualExpendRequest
}

type AssetAndLiabilityData struct {
	TotalLiquidAssets         int64 `json:"total_liquid_assets"`          // Total Aset Likuid
	TotalInvestmentAssets     int64 `json:"total_investment_assets"`      // Total
	TotalPersonalAssets       int64 `json:"total_personal_assets"`        // Total Aset Personal
	TotalAssets               int64 `json:"total_assets"`                 // Total Aset
	TotalShortTermLiabilities int64 `json:"total_short_term_liabilities"` //
	TotalLongTermLiabilities  int64 `json:"total_long_term_liabilities"`  // Total Kewajiban Jangka Panjang
	TotalLiabilities          int64 `json:"total_liabilities"`            // Total Kew
	TotalNetWorth             int64 `json:"total_net_worth"`              // Total Kekayaan Bersih
}

type KasFlowData struct {
	TotalActiveIncome  int64 `json:"total_active_income"`  // Total Pendapatan Aktif
	TotalPassiveIncome int64 `json:"total_passive_income"` // Total Pendapatan Pasif
	TotalIncome        int64 `json:"total_income"`         // Total Pendapatan
}

type AnnualExpendData struct {
	TotalPosSosial                      int64 `json:"total_pos_sosial"`                         // Total Pos Sosial
	TotalPosTabunganInvestasi           int64 `json:"total_pos_tabungan_investasi"`             // Total Pos Tabungan & Investasi
	TotalPosBayarKewajibanJangkaPendek  int64 `json:"total_pos_bayar_kewajiban_jangka_pendek"`  // Total Pos Bayar Kewajiban Jangka Pendek
	TotalPosBayarKewajibanJangkaPanjang int64 `json:"total_pos_bayar_kewajiban_jangka_panjang"` // Total Pos Bayar Kewaj
	TotalPosBayarPremiAsuransi          int64 `json:"total_pos_bayar_premi_asuransi"`           // Total Pos Bayar Premi Asuransi
	TotalPosBiayaHidupMasaKini          int64 `json:"total_pos_biaya_hidup_masa_kini"`          // Total Pos Biaya Hidup Masa Kini
	TotalPengeluaran                    int64 `json:"total_pengeluaran"`                        // Total Pengeluaran Tahunan
	TotalPendapatan                     int64 `json:"total_pendapatan"`                         // Total Pendapatan Bersih Tahunan
	ArusKas                             int64 `json:"arus_kas"`                                 // Arus Kas Bersih Tahunan
}

type FinancialCheckUpData struct {
	AssetAndLiabilityData AssetAndLiabilityData `json:"asset_and_liability_data"` // Data Aset dan Kewajiban
	KasFlowData           KasFlowData           `json:"kas_flow_data"`            // Data Kas Flow
	AnnualExpendData      AnnualExpendData      `json:"annual_expend_data"`       // Data Pengeluaran Tahunan
}

type AssetAndLiabilityResponse struct {
	Status int                   `json:"status"` // Status dari permintaan
	Data   AssetAndLiabilityData `json:"data"`   // Data hasil perhitungan
}

type KasFlowResponse struct {
	Status int         `json:"status"` // Status dari permintaan
	Data   KasFlowData `json:"data"`   // Data hasil perhitungan
}

type AnnualExpendResponse struct {
	Status int              `json:"status"` // Status dari permintaan
	Data   AnnualExpendData `json:"data"`   // Data hasil perhitungan
}

type FinancialCheckUpResponse struct {
	Status int                  `json:"status"` // Status dari permintaan
	Data   FinancialCheckUpData `json:"data"`   // Data hasil perhitungan
}
