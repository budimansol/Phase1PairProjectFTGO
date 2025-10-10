package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type ReportService interface {
	StaffReport(startDate, endDate string) ([]model.StaffReport, error)
	OrderReport(startDate, endDate string) ([]model.OrderReport, error)
	StockReport(startDate, endDate string) ([]model.StockReport, error)
	DailySalesReport(startDate, endDate string) ([]model.SalesSummary, error)
	MonthlySalesReport(startDate, endDate string) ([]model.SalesSummary, error)
	SalesByCategory() ([]model.SalesByCategory, error)
	AboveAverageMember() ([]model.AboveAverageMember, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (r *reportService) StaffReport(startDate, endDate string) ([]model.StaffReport, error) {
	return r.repo.GetStaffReport(startDate, endDate)
}

func (r *reportService) OrderReport(startDate, endDate string) ([]model.OrderReport, error) {
	return r.repo.GetOrderReport(startDate, endDate)
}

func (r *reportService) StockReport(startDate, endDate string) ([]model.StockReport, error) {
	return r.repo.GetStockReport(startDate, endDate)
}

func (r *reportService) DailySalesReport(startDate, endDate string) ([]model.SalesSummary, error) {
	return r.repo.GetDailySalesReport(startDate, endDate)
}

func (r *reportService) MonthlySalesReport(startDate, endDate string) ([]model.SalesSummary, error) {
	return r.repo.GetMonthlySalesReport(startDate, endDate)
}

func (r *reportService) AboveAverageMember() ([]model.AboveAverageMember, error) {
	return r.repo.GetAboveAverageMember()
}

func (r *reportService) SalesByCategory() ([]model.SalesByCategory, error) {
	return r.repo.GetSalesByCategory()
}
