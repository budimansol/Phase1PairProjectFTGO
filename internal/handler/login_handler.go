package handler

import (
	"fmt"

	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type LoginHandler struct {
	staffService *service.StaffService
}

func NewLoginHandler(staffService *service.StaffService) *LoginHandler {
	return &LoginHandler{staffService}
}

func (h *LoginHandler) Login() (*string, error) {
	for {
		fmt.Println("🔐 Staff Login")

		// Prompt email
		emailPrompt := promptui.Prompt{
			Label: "Email",
		}
		email, _ := emailPrompt.Run()

		// Prompt password (hidden input)
		passPrompt := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}
		password, _ := passPrompt.Run()

		staff, err := h.staffService.Login(email, password)
		if err != nil {
			fmt.Println("❌ Login failed:", err)
			continue
		}

		fmt.Printf("✅ Welcome, %s!\n", staff.Name)
		return &staff.Name, nil
	}
}
