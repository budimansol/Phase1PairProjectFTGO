package repository

import (
	"database/sql"

	"github.com/budimansol/pairproject/internal/model"
)

type MemberRepository interface {
	CreateMember(member *model.Member) error
	UpdateMember(member *model.Member, id int) error
	DeleteMember(id int) error
	GetMemberByID(id int) (*model.Member, error)
	GetAllMembers() ([]model.Member, error)
}

type memberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (repo *memberRepository) CreateMember(member *model.Member) error {
	_, err := repo.db.Exec(
		`INSERT INTO members (name, email, phone, reward_points) VALUES ($1, $2, $3, $4)`,
		member.Name, member.Email, member.Phone, member.RewardPoints,
	)
	return err
}

func (repo *memberRepository) UpdateMember(member *model.Member, id int) error {
	_, err := repo.db.Exec(`UPDATE members SET name = $1, phone = $2, email = $3, reward_points = $4 WHERE id = $5`, member.Name, member.Phone, member.Email, member.RewardPoints, id)
	if err != nil {
		return err
	}
	return nil

}

func (repo *memberRepository) DeleteMember(id int) error {
	_, err := repo.db.Exec(`DELETE FROM members WHERE id = $1`, id)
	return err
}

func (repo *memberRepository) GetMemberByID(id int) (*model.Member, error) {
	var member model.Member
	err := repo.db.QueryRow("SELECT id, name, phone, email, reward_points FROM members WHERE id = ?", id).Scan(&member.ID, &member.Name, &member.Phone, &member.Email, &member.RewardPoints)
	return &member, err
}

func (repo *memberRepository) GetAllMembers() ([]model.Member, error) {
	rows, err := repo.db.Query("SELECT id, name, phone, email, reward_points FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.Member
	for rows.Next() {
		var member model.Member
		err := rows.Scan(&member.ID, &member.Name, &member.Phone, &member.Email, &member.RewardPoints)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}
