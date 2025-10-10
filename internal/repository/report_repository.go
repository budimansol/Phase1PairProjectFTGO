package repository

import (
	"database/sql"

	"github.com/budimansol/pairproject/internal/model"
)

type ReportRepository interface {
	GetStaffReport(startDate, endDate string) ([]model.StaffReport, error)
	GetOrderReport(startDate, endDate string) ([]model.OrderReport, error)
	GetStockReport(startDate, endDate string) ([]model.StockReport, error)
	GetDailySalesReport(startDate, endDate string) ([]model.SalesSummary, error)
	GetMonthlySalesReport(startDate, endDate string) ([]model.SalesSummary, error)
	GetSalesByCategory() ([]model.SalesByCategory, error)
	GetAboveAverageMember() ([]model.AboveAverageMember, error)
}

type reportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{DB: db}
}

func (r *reportRepository) GetStaffReport(startDate, endDate string) ([]model.StaffReport, error) {
	query := `
	WITH transaction_summary AS (
    	SELECT 
        staff_id,
        SUM(total_amount) AS total_sales
    	FROM transactions
    	WHERE transaction_date BETWEEN $1 AND $2
    	GROUP BY staff_id
	)
	SELECT 
    s.id AS staff_id,
    s.name AS staff_name,
    COALESCE(ts.total_sales, 0) AS total_sales
	FROM staffs s
	LEFT JOIN transaction_summary ts ON s.id = ts.staff_id
	ORDER BY total_sales DESC;`

	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.StaffReport
	for rows.Next() {
		var rep model.StaffReport
		if err := rows.Scan(&rep.StaffID, &rep.StaffName, &rep.Total); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}

	return reports, nil
}

func (r *reportRepository) GetOrderReport(startDate, endDate string) ([]model.OrderReport, error) {
	query := `WITH sales AS (
        SELECT 
            ti.menu_id,
            SUM(ti.quantity) AS total_sold,
            SUM(ti.subtotal) AS total_revenue
        FROM transaction_items ti
        JOIN transactions t ON ti.transaction_id = t.id
        WHERE t.transaction_date BETWEEN $1 AND $2
        GROUP BY ti.menu_id
    )
    SELECT 
        m.id, m.name, m.category,
        COALESCE(s.total_sold, 0),
        COALESCE(s.total_revenue, 0)
    FROM menus m
    LEFT JOIN sales s ON m.id = s.menu_id
    ORDER BY s.total_revenue DESC;`

	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.OrderReport
	for rows.Next() {
		var rep model.OrderReport
		if err := rows.Scan(&rep.MenuID, &rep.MenuName, &rep.Category, &rep.TotalSold, &rep.TotalRevenue); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *reportRepository) GetStockReport(startDate, endDate string) ([]model.StockReport, error) {
	query := `
    WITH sold_items AS (
        SELECT 
            ti.menu_id,
            SUM(ti.quantity) AS total_sold
        FROM transaction_items ti
        JOIN transactions t ON ti.transaction_id = t.id
        WHERE t.transaction_date BETWEEN $1 AND $2
        GROUP BY ti.menu_id
    )
    SELECT 
        m.id AS menu_id,
        m.name AS menu_name,
        m.category,
        m.stock AS current_stock,
        COALESCE(s.total_sold, 0) AS total_sold,
        (m.stock * m.price) AS stock_value
    FROM menus m
    LEFT JOIN sold_items s ON m.id = s.menu_id
    ORDER BY m.stock ASC;
    `
	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.StockReport
	for rows.Next() {
		var rep model.StockReport
		if err := rows.Scan(
			&rep.MenuID,
			&rep.MenuName,
			&rep.Category,
			&rep.CurrentStock,
			&rep.TotalSold,
			&rep.StockValue,
		); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *reportRepository) GetDailySalesReport(startDate, endDate string) ([]model.SalesSummary, error) {
	query := `
		SELECT 
			TO_CHAR(transaction_date, 'YYYY-MM-DD') AS day,
			COUNT(id) AS total_order,
			COALESCE(SUM(total_amount), 0) AS total_sales
		FROM transactions
		WHERE transaction_date BETWEEN $1 AND $2
		GROUP BY day
		ORDER BY day ASC;
	`
	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.SalesSummary
	for rows.Next() {
		var rep model.SalesSummary
		if err := rows.Scan(&rep.Period, &rep.TotalOrder, &rep.TotalSales); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *reportRepository) GetMonthlySalesReport(startDate, endDate string) ([]model.SalesSummary, error) {
	query := `
		SELECT 
			TO_CHAR(transaction_date, 'YYYY-MM') AS month,
			COUNT(id) AS total_order,
			COALESCE(SUM(total_amount), 0) AS total_sales
		FROM transactions
		WHERE transaction_date BETWEEN $1 AND $2
		GROUP BY month
		ORDER BY month ASC;
	`
	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.SalesSummary
	for rows.Next() {
		var rep model.SalesSummary
		if err := rows.Scan(&rep.Period, &rep.TotalOrder, &rep.TotalSales); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *reportRepository) GetSalesByCategory() ([]model.SalesByCategory, error) {
	query := `
		SELECT 
			COALESCE(m.category, 'TOTAL') AS category,
			COALESCE(SUM(ti.subtotal), 0) AS total_sales
		FROM transaction_items ti
		JOIN menus m ON ti.menu_id = m.id
		GROUP BY ROLLUP(m.category);
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.SalesByCategory
	for rows.Next() {
		var rep model.SalesByCategory
		if err := rows.Scan(&rep.Category, &rep.TotalSales); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *reportRepository) GetAboveAverageMember() ([]model.AboveAverageMember, error) {
	query := `
		SELECT m.id id, m.name name,  (
		SELECT AVG(total_amount)
		FROM transactions t
		WHERE t.member_id = m.id
		) as average,
		( SELECT AVG(total_amount) FROM transactions) as avg_all
		FROM members m
		WHERE (
		SELECT AVG(total_amount)
		FROM transactions t
		WHERE t.member_id = m.id
		) > (
		SELECT AVG(total_amount) FROM transactions);`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.AboveAverageMember
	for rows.Next() {
		var rep model.AboveAverageMember
		if err := rows.Scan(&rep.ID, &rep.Name, &rep.Average, &rep.AvgAll); err != nil {
			return nil, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}
