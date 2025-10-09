package model

type Transaction struct {
	ID              int
	StaffID         int
	MemberID        int
	TransactionDate string
	TotalAmount     float64
}
