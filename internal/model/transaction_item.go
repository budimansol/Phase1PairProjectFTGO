package model

type TransactionItem struct {
	ID            int
	TransactionID int
	MenuID        int
	MenuName      string
	Price         float64
	Quantity      int
	Subtotal      float64
}