package handler

import (
	"fmt"
	"strconv"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type StaffHandler struct {
	service *service.StaffService
}

func NewStaffHandler(s *service.StaffService) *StaffHandler {
	return &StaffHandler{service: s}
}

func (h *StaffHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Staff Management",
			Items: []string{
				"1. Create", 
				"2. List", 
				"3. Update", 
				"4. Delete", 
				"5. Exit"},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Create":
			h.create()
		case "2. List":
			h.list()
		case "3. Update":
			h.update()
		case "4. Delete":
			h.delete()
		case "5. Exit":
			return
		}
	}
}

func (h *StaffHandler) create() {
	namePrompt := promptui.Prompt{Label: "Name"}
	emailPrompt := promptui.Prompt{Label: "Email"}
	rolePrompt := promptui.Prompt{Label: "Role"}
	passPrompt := promptui.Prompt{Label: "Password"}

	name, _ := namePrompt.Run()
	email, _ := emailPrompt.Run()
	role, _ := rolePrompt.Run()
	pass, _ := passPrompt.Run()

	staff := model.Staff{Name: name, Email: email, Role: role, Password: pass}
	err := h.service.CreateStaff(staff)
	if err != nil {
		fmt.Println("❌ Failed to create staff:", err)
		return
	}
	fmt.Println("✅ Staff created successfully!")
}

func (h *StaffHandler) list() {
	staffs, err := h.service.GetAllStaff()
	if err != nil {
		fmt.Println("❌ Error fetching staff:", err)
		return
	}
	fmt.Println("\n=== Staff List ===")
	for _, s := range staffs {
		fmt.Printf("%d. %s (%s) - %s\n", s.ID, s.Name, s.Email, s.Role)
	}
}

func (h *StaffHandler) update() {
	idPrompt := promptui.Prompt{Label: "Staff ID"}
	idStr, _ := idPrompt.Run()
	id, _ := strconv.Atoi(idStr)

	namePrompt := promptui.Prompt{Label: "New Name"}
	emailPrompt := promptui.Prompt{Label: "New Email"}
	rolePrompt := promptui.Prompt{Label: "New Role"}

	name, _ := namePrompt.Run()
	email, _ := emailPrompt.Run()
	role, _ := rolePrompt.Run()

	staff := model.Staff{ID: id, Name: name, Email: email, Role: role}
	err := h.service.UpdateStaff(staff)
	if err != nil {
		fmt.Println("❌ Update failed:", err)
		return
	}
	fmt.Println("✅ Staff updated successfully!")
}

func (h *StaffHandler) delete() {
	idPrompt := promptui.Prompt{Label: "Staff ID"}
	idStr, _ := idPrompt.Run()
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteStaff(id)
	if err != nil {
		fmt.Println("❌ Failed to delete staff:", err)
		return
	}
	fmt.Println("✅ Staff deleted successfully!")
}
