package handler

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type MainHandler struct {
	StaffHandler  *StaffHandler
	MemberHandler *MemberHandler
	ReportHandler *ReportHandler
	MenuHandler	  *MenuHandler
	ReservationHandler *ReservationHandler
}

func NewMainHandler(
	staffHandler *StaffHandler,
	memberHandler *MemberHandler,
	reportHandler *ReportHandler,
  menuHandler *MenuHandler,
  reservationHandler *ReservationHandler
) *MainHandler {
	return &MainHandler{
		StaffHandler:  staffHandler,
		MemberHandler: memberHandler,
		ReportHandler: reportHandler,
    MenuHandler: menuHandler,
    ReservationHandler : reservationHandler
)

func (h *MainHandler) Run() {
	for {
		prompt := promptui.Select{
			Label: "🏪 Beverage CLI Main Menu",
			Items: []string{
				"1. Staff Management",
				"2. Menu Management",
				"3. Reservation Management",
        "4. Member Management",
				"5. Reports",
				"6. Exit",
			},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Staff Management":
			h.StaffHandler.Menu()
    case "2. Menu Management":
			h.MenuHandler.Menu()
    case "3. Reservation Management":
			h.ReservationHandler.Menu()
		case "4. Member Management":
			h.MemberHandler.Menu()
		case "5. Reports":
			h.ReportHandler.Menu()
		case "6. Exit":
			fmt.Println("👋 Exiting Beverage CLI... Goodbye!")
			return
		}
	}
}
