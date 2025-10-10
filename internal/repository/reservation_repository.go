package repository

import (
	"database/sql"

	"github.com/budimansol/pairproject/internal/model"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

// CREATE
func (r *ReservationRepository) Create(res model.Reservation) error {
	query := `
		INSERT INTO reservations (member_id, reservation_date, time_slot, total_people, note)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query, res.MemberID, res.ReservationDate, res.TimeSlot, res.TotalPeople, res.Note)
	return err
}

// GET ALL
func (r *ReservationRepository) GetAll() ([]model.Reservation, error) {
	rows, err := r.db.Query(`
		SELECT id, member_id, reservation_date, time_slot, total_people, note, created_at
		FROM reservations
		ORDER BY reservation_date, time_slot
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []model.Reservation
	for rows.Next() {
		var rsv model.Reservation
		err := rows.Scan(
			&rsv.ID, &rsv.MemberID, &rsv.ReservationDate,
			&rsv.TimeSlot, &rsv.TotalPeople, &rsv.Note, &rsv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, rsv)
	}
	return reservations, nil
}
