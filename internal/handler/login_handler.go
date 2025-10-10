package handler

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type LoginHandler struct {
	staffService *service.StaffService
}

func NewLoginHandler(staffService *service.StaffService) *LoginHandler {
	return &LoginHandler{staffService}
}

func (h *LoginHandler) Login() (*model.Staff, error) {
	emailPrompt := promptui.Prompt{Label: "Email"}
	email, _ := emailPrompt.Run()

	passwordPrompt := promptui.Prompt{Label: "Password", Mask: '*'}
	password, _ := passwordPrompt.Run()

	return h.staffService.Login(email, password)
}
