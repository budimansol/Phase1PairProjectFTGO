package service

import (
	"fmt"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type StaffService struct {
	repo *repository.StaffRepository
}

func NewStaffService(repo *repository.StaffRepository) *StaffService {
	return &StaffService{repo: repo}
}

func (s *StaffService) CreateStaff(staff model.Staff) error {
	return s.repo.Create(staff)
}

func (s *StaffService) GetAllStaff() ([]model.Staff, error) {
	return s.repo.GetAll()
}

func (s *StaffService) UpdateStaff(staff model.Staff) error {
	return s.repo.Update(staff)
}

func (s *StaffService) DeleteStaff(id int) error {
	return s.repo.Delete(id)
}

func (s *StaffService) Login(email, password string) (*model.Staff, error) {
	staff, err := s.repo.FindByEmail(email)
	if err != nil || staff == nil {
		return nil, fmt.Errorf("email atau password salah")
	}

	if staff.Password != password {
		return nil, fmt.Errorf("email atau password salah")
	}

	return staff, nil
}
