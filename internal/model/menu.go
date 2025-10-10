package model

import "time"

type Menu struct {
	ID        int
	Name      string
	Category  string
	Price     float64
	Stock     int
	CreatedAt time.Time
}
