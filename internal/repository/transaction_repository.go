package repository

import (
	"database/sql"
	"github.com/budimansol/pairproject/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create transaction utama
func (r *TransactionRepository) Create(tx *model.Transaction) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO transactions (staff_id, member_id, transaction_date, total_amount)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, tx.StaffID, tx.MemberID, tx.TransactionDate, tx.TotalAmount).Scan(&id)
	return id, err
}

// Tambahkan item transaksi
func (r *TransactionRepository) AddItem(item *model.TransactionItem) error {
	_, err := r.db.Exec(`
		INSERT INTO transaction_items (transaction_id, menu_id, quantity, subtotal)
		VALUES ($1, $2, $3, $4)
	`, item.TransactionID, item.MenuID, item.Quantity, item.Subtotal)
	return err
}

// Ambil semua transaksi beserta item
func (r *TransactionRepository) GetAll() ([]model.Transaction, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.staff_id, s.name, t.member_id, m.name, t.transaction_date, t.total_amount
		FROM transactions t
		JOIN staffs s ON t.staff_id = s.id
		LEFT JOIN members m ON t.member_id = m.id
		ORDER BY t.transaction_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		var memberID sql.NullInt64
		var memberName sql.NullString
		var staffName string

		if err := rows.Scan(&tx.ID, &tx.StaffID, &staffName, &memberID, &memberName, &tx.TransactionDate, &tx.TotalAmount); err != nil {
			return nil, err
		}
		tx.StaffName = staffName
		if memberID.Valid {
			id := int(memberID.Int64)
			tx.MemberID = &id
		}
		if memberName.Valid {
			tx.MemberName = memberName.String
		} else {
			tx.MemberName = "Guest"
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

// Hapus transaksi beserta item
func (r *TransactionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM transaction_items WHERE transaction_id = $1", id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}
