package model

import "time"

type StaffProfile struct {
	ID               int
	StaffID          int
	Address          string
	Phone            string
	DateOfBirth      time.Time
	EmergencyContact string
}
