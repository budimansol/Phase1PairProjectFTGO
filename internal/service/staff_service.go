package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type StaffService interface {
	CreateStaff(staff *model.Staff) error
}

type staffService struct {
	repo repository.StaffRepository
}
