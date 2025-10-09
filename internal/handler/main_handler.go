package handler

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type MainHandler struct {
	StaffHandler  *StaffHandler
}

func NewMainHandler(
	staffHandler *StaffHandler,
) *MainHandler {
	return &MainHandler{
		StaffHandler:  staffHandler,
	}
}

func (h *MainHandler) Run() {
	for {
		prompt := promptui.Select{
			Label: "🏪 Beverage CLI Main Menu",
			Items: []string{
				"1. Staff Management",
				"2. Exit",
			},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Staff Management":
			h.StaffHandler.Menu()
		case "2. Exit":
			fmt.Println("👋 Exiting Beverage CLI... Goodbye!")
			return
		}
	}
}
