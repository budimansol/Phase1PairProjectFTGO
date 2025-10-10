package model

type StaffReport struct {
	StaffID   int
	StaffName string
	Total     int
}

type OrderReport struct {
	MenuID       int
	MenuName     string
	Category     string
	TotalSold    int
	TotalRevenue float64
}

type StockReport struct {
	MenuID       int
	MenuName     string
	Category     string
	CurrentStock int
	TotalSold    int
	StockValue   float64
}

type SalesSummary struct {
	Period     string
	TotalSales float64
	TotalOrder int
}

type SalesByCategory struct {
	Category   string
	TotalSales float64
}

type AboveAverageMember struct {
	ID      int
	Name    string
	Average float64
	AvgAll  float64
}
