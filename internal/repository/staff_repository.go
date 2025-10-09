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
	query := `
	SELECT 
		s.id, s.name, s.email, s.role, s.created_at,
		p.id, p.staff_id, p.address, p.phone, p.date_of_birth, p.emergency_contact
	FROM staffs s
	LEFT JOIN staff_profiles p ON s.id = p.staff_id
	ORDER BY s.id;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []model.Staff

	for rows.Next() {
		var s model.Staff
		var p model.StaffProfile
		var dob sql.NullTime
		var addr, phone, emergency sql.NullString

		err := rows.Scan(
			&s.ID, &s.Name, &s.Email, &s.Role, &s.CreatedAt,
			&p.ID, &p.StaffID, &addr, &phone, &dob, &emergency, 
		)
		if err != nil {
			return nil, err
		}

		if addr.Valid || phone.Valid || dob.Valid || emergency.Valid {
			s.Profile = &model.StaffProfile{
				ID:               p.ID,
				StaffID:          p.StaffID,
				Address:          addr.String,
				Phone:            phone.String,
				DateOfBirth:      dob.Time,
				EmergencyContact: emergency.String,
			}
		}

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

func (r *StaffRepository) FindByEmail(email string) (*model.Staff, error) {
	query := "SELECT id, name, email, password, role FROM staffs WHERE email = $1"
	row := r.DB.QueryRow(query, email)

	var s model.Staff
	err := row.Scan(&s.ID, &s.Name, &s.Email, &s.Password, &s.Role)
	if err != nil {
		return nil, err
	}
	return &s, nil
}