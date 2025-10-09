package repository

import (
	"database/sql"

	"github.com/budimansol/pairproject/internal/model"
)

type StaffRepository struct {
	DB *sql.DB
}

func NewStaffRepository(db *sql.DB) *StaffRepository {
	return &StaffRepository{DB: db}
}

func (r *StaffRepository) Create(staff model.Staff) error {
	_, err := r.DB.Exec(
		`INSERT INTO staffs (name, email, password, role) VALUES ($1, $2, $3, $4)`,
		staff.Name, staff.Email, staff.Password, staff.Role,
	)
	return err
}

func (r *StaffRepository) GetAll() ([]model.Staff, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, role, created_at FROM staffs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []model.Staff
	for rows.Next() {
		var s model.Staff
		rows.Scan(&s.ID, &s.Name, &s.Email, &s.Role, &s.CreatedAt)
		staffs = append(staffs, s)
	}
	return staffs, nil
}

func (r *StaffRepository) Update(staff model.Staff) error {
	_, err := r.DB.Exec(`UPDATE staffs SET name=$1, email=$2, role=$3 WHERE id=$4`,
		staff.Name, staff.Email, staff.Role, staff.ID)
	return err
}

func (r *StaffRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM staffs WHERE id=$1`, id)
	return err
}
