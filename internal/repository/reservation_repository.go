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
		SELECT r.id, r.member_id, m.name, r.reservation_date, r.time_slot, r.total_people, r.note, r.created_at
		FROM reservations r
		JOIN members m ON r.member_id = m.id
		ORDER BY r.reservation_date, r.time_slot
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []model.Reservation
	for rows.Next() {
		var res model.Reservation
		if err := rows.Scan(&res.ID, &res.MemberID, &res.MemberName, &res.ReservationDate, &res.TimeSlot, &res.TotalPeople, &res.Note, &res.CreatedAt); err != nil {
			return nil, err
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}

func (r *ReservationRepository) Update(reservation *model.Reservation) error {
	_, err := r.db.Exec(`
		UPDATE reservations 
		SET member_id = $1, reservation_date = $2, time_slot = $3, total_people = $4, note = $5
		WHERE id = $6
	`, reservation.MemberID, reservation.ReservationDate, reservation.TimeSlot, reservation.TotalPeople, reservation.Note, reservation.ID)
	return err
}


func (r *ReservationRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM reservations WHERE id = $1`, id)
	return err
}
