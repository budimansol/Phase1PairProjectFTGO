package model

import "time"

type Reservation struct {
	ID             int
	MemberID       int
	ReservationDate time.Time
	TimeSlot       string
	TotalPeople    int
	Note           string
	CreatedAt      time.Time
}
