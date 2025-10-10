package handler

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type MainHandler struct {
	StaffHandler  *StaffHandler
	MemberHandler *MemberHandler
	ReportHandler *ReportHandler
}

func NewMainHandler(
	staffHandler *StaffHandler,
	memberHandler *MemberHandler,
	reportHandler *ReportHandler,
) *MainHandler {
	return &MainHandler{
		StaffHandler:  staffHandler,
		MemberHandler: memberHandler,
		ReportHandler: reportHandler,
	}
}

func (h *MainHandler) Run() {
	for {
		prompt := promptui.Select{
			Label: "🏪 Beverage CLI Main Menu",
			Items: []string{
				"1. Staff Management",
				"6. Member Management",
				"7. Reports",
				"9. Exit",
			},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Staff Management":
			h.StaffHandler.Menu()
		case "6. Member Management":
			h.MemberHandler.Menu()
		case "7. Reports":
			h.ReportHandler.Menu()
		case "9. Exit":
			fmt.Println("👋 Exiting Beverage CLI... Goodbye!")
			return
		}
	}
}
