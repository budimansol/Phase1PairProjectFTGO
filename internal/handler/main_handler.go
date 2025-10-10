package handler

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type MainHandler struct {
	StaffHandler  *StaffHandler
	MenuHandler	  *MenuHandler
}

func NewMainHandler(
	staffHandler *StaffHandler,
	menuHandler *MenuHandler,
) *MainHandler {
	return &MainHandler{
		StaffHandler:  staffHandler,
		MenuHandler: menuHandler,
	}
}

func (h *MainHandler) Run() {
	for {
		prompt := promptui.Select{
			Label: "🏪 Beverage CLI Main Menu",
			Items: []string{
				"1. Staff Management",
				"2. Menu Management",
				"3. Exit",
			},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Staff Management":
			h.StaffHandler.Menu()
		case "2. Menu Management":
			h.MenuHandler.Menu()
		case "3. Exit":
			fmt.Println("👋 Exiting Beverage CLI... Goodbye!")
			return
		}
	}
}
