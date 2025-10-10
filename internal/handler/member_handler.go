package handler

import (
	"fmt"
	"strconv"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type MemberHandler struct {
	service service.MemberService
}

func NewMemberHandler(s service.MemberService) *MemberHandler {
	return &MemberHandler{service: s}
}

func (h *MemberHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Member Management",
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

func (h *MemberHandler) create() {
	namePrompt := promptui.Prompt{Label: "Name"}
	phonePrompt := promptui.Prompt{Label: "Phone"}
	emailPrompt := promptui.Prompt{Label: "Email"}

	name, _ := namePrompt.Run()
	phone, _ := phonePrompt.Run()
	email, _ := emailPrompt.Run()

	member := model.Member{Name: name, Phone: phone, Email: email, RewardPoints: 0}
	err := h.service.CreateMember(&member)
	if err != nil {
		fmt.Println("❌ Failed to create member:", err)
		return
	}
	fmt.Println("✅ Member created successfully!")
}

func (h *MemberHandler) list() {
	members, err := h.service.GetAllMembers()
	if err != nil {
		fmt.Println("❌ Error fetching members:", err)
		return
	}
	fmt.Println("\n=== Member List ===")
	for _, s := range members {
		fmt.Printf("%d. %s (%s) - %s - %d\n", s.ID, s.Name, s.Email, s.Phone, s.RewardPoints)
	}
}

func (h *MemberHandler) update() {
	idPrompt := promptui.Prompt{Label: "Member ID"}
	idStr, _ := idPrompt.Run()
	id, _ := strconv.Atoi(idStr)

	namePrompt := promptui.Prompt{Label: "New Name"}
	emailPrompt := promptui.Prompt{Label: "New Email"}
	phonePrompt := promptui.Prompt{Label: "New Phone"}

	name, _ := namePrompt.Run()
	email, _ := emailPrompt.Run()
	phone, _ := phonePrompt.Run()

	member := model.Member{ID: id, Name: name, Email: email, Phone: phone}
	err := h.service.UpdateMember(&member, id)
	if err != nil {
		fmt.Println("❌ Update failed:", err)
		return
	}
	fmt.Println("✅ Staff updated successfully!")
}

func (h *MemberHandler) delete() {
	idPrompt := promptui.Prompt{Label: "Member ID"}
	idStr, _ := idPrompt.Run()
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteMember(id)
	if err != nil {
		fmt.Println("❌ Failed to delete staff:", err)
		return
	}
	fmt.Println("✅ Staff deleted successfully!")
}
