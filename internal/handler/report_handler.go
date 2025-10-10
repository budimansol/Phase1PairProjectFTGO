package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(s service.ReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

func (h *ReportHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Reports Menu",
			Items: []string{
				"1. Sales By Staff",
				"2. Sales By Menu",
				"3. Sales By Category",
				"4. Daily Sales Report",
				"5. Monthly Sales Report",
				"6. Priority Member",
				"7. Stock Report",
				"8. Exit"},
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . | cyan }}:",
				Active:   "> {{ . | green }}",
				Inactive: "  {{ . | faint }}",
				Selected: "✔ {{ . | bold }}",
			},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Sales By Staff":
			h.staffReport()
		case "2. Sales By Menu":
			h.orderReport()
		case "3. Sales By Category":
			h.salesByCategory()
		case "4. Daily Sales Report":
			h.dailyReport()
		case "5. Monthly Sales Report":
			h.monthlyReport()
		case "6. Priority Member":
			h.aboveAverageMember()
		case "7. Stock Report":
			h.stockReport()
		case "8. Exit":
			return
		}
	}
}

func validateDateInput(label string) (string, error) {
	for {
		prompt := promptui.Prompt{Label: label}
		dateStr, err := prompt.Run()
		if err != nil {
			return "", err
		}
		_, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("[ x ] Format tanggal salah! Gunakan format YYYY-MM-DD.")
			continue
		}
		return dateStr, nil
	}
}

func (h *ReportHandler) staffReport() {
	startDate, _ := validateDateInput("Start Date (YYYY-MM-DD)")
	endDate, _ := validateDateInput("End Date (YYYY-MM-DD)")

	staffReport, err := h.service.StaffReport(startDate, endDate)
	if err != nil {
		fmt.Println("[ x ] Failed to generate staff report:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name", "Total Sales"})
	for _, report := range staffReport {
		table.Append([]string{
			fmt.Sprintf("%d", report.StaffID),
			report.StaffName,
			fmt.Sprintf("%d", report.Total),
		})
	}
	fmt.Println("\n=== Staff Report ===")
	table.Render()
}

func (h *ReportHandler) orderReport() {
	startDate, _ := validateDateInput("Start Date (YYYY-MM-DD)")
	endDate, _ := validateDateInput("End Date (YYYY-MM-DD)")

	reports, err := h.service.OrderReport(startDate, endDate)
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Menu Name", "Qty Sold", "Total Revenue"})
	for _, r := range reports {

		table.Append([]string{
			fmt.Sprintf("%v", r.MenuName),
			fmt.Sprintf("%v", r.TotalSold),
			fmt.Sprintf("Rp %.2f", r.TotalRevenue),
		})
	}
	fmt.Println("\n=== Order Report ===")
	table.Render()
}

func (h *ReportHandler) dailyReport() {
	startDate, _ := validateDateInput("Start Date (YYYY-MM-DD)")
	endDate, _ := validateDateInput("End Date (YYYY-MM-DD)")

	reports, err := h.service.DailySalesReport(startDate, endDate)
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Date", "Total Orders", "Total Sales"})
	for _, r := range reports {
		table.Append([]string{
			r.Period,
			fmt.Sprintf("%d", r.TotalOrder),
			fmt.Sprintf("%.2f", r.TotalSales),
		})
	}
	fmt.Println("\n=== Daily Sales Report ===")
	table.Render()
}

func (h *ReportHandler) monthlyReport() {
	startDate, _ := validateDateInput("Start Date (YYYY-MM-DD)")
	endDate, _ := validateDateInput("End Date (YYYY-MM-DD)")

	reports, err := h.service.MonthlySalesReport(startDate, endDate)
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Month", "Total Orders", "Total Sales"})
	for _, r := range reports {
		table.Append([]string{
			r.Period,
			fmt.Sprintf("%d", r.TotalOrder),
			fmt.Sprintf("%.2f", r.TotalSales),
		})
	}
	fmt.Println("\n=== Monthly Sales Report ===")
	table.Render()
}

func (h *ReportHandler) stockReport() {
	startDate, _ := validateDateInput("Start Date (YYYY-MM-DD)")
	endDate, _ := validateDateInput("End Date (YYYY-MM-DD)")

	reports, err := h.service.StockReport(startDate, endDate)
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Menu Name", "Category", "Stock", "Sold", "Stock Value"})
	for _, r := range reports {
		table.Append([]string{
			fmt.Sprintf("%v", r.MenuName),
			fmt.Sprintf("%v", r.Category),
			fmt.Sprintf("%v", r.CurrentStock),
			fmt.Sprintf("%v", r.TotalSold),
			fmt.Sprintf("Rp %.2f", r.StockValue),
		})
	}
	fmt.Println("\n=== Stock Report ===")
	table.Render()
}

func (h *ReportHandler) salesByCategory() {
	reports, err := h.service.SalesByCategory()
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Category", "Total Sales"})
	for _, r := range reports {
		table.Append([]string{
			fmt.Sprintf("%v", r.Category),
			fmt.Sprintf("%v", r.TotalSales),
		})
	}
	fmt.Println("\n=== Sales By Category ===")
	table.Render()
}

func (h *ReportHandler) aboveAverageMember() {
	reports, err := h.service.AboveAverageMember()
	if err != nil {
		fmt.Println("[ x ] Error:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name", "Average Member", "Average All"})
	for _, r := range reports {
		table.Append([]string{
			fmt.Sprintf("%v", r.ID),
			fmt.Sprintf("%v", r.Name),
			fmt.Sprintf("%v", r.Average),
			fmt.Sprintf("%v", r.AvgAll),
		})
	}
	fmt.Println("\n=== Priority Member ===")
	table.Render()
}
