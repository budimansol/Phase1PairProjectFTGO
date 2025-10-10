package model

import "time"

type Transaction struct {
	ID              int                // transaction ID
	StaffID         int                // FK staff
	StaffName       string             // nama staff (untuk tampilan)
	MemberID        *int               // FK member, nullable
	MemberName      string             // nama member (untuk tampilan)
	TransactionDate time.Time          // timestamp
	TotalAmount     float64            // total
	Items           []TransactionItem  // daftar item transaksi
}