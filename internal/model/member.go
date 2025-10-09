package model

import "time"

type Member struct {
	ID           int
	Name         string
	Phone        string
	Email        string
	RewardPoints int
	CreatedAt    time.Time
}
