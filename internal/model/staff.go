package model

import "time"

type Staff struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
	CreatedAt time.Time
}
