package handler

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
)

type StaffHandler struct {
	StaffService *service.StaffService
}

func NewStaffHandler(staffService *service.StaffService) *StaffHandler {
	return &StaffHandler{StaffService: staffService}
}

func (h *StaffHandler) CreateStaff(staff *model.Staff) error {
	return h.StaffService.CreateStaff(staff)
}
