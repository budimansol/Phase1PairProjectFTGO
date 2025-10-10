package handler

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type MainHandler struct {
	LoginHandler *LoginHandler
	StaffHandler  *StaffHandler
	MenuHandler	  *MenuHandler
	ReservationHandler *ReservationHandler
	TransactionHandler *TransactionHandler
}

func NewMainHandler(
	loginHandler *LoginHandler,
	staffHandler *StaffHandler,
	menuHandler *MenuHandler,
	reservationHandler *ReservationHandler,
	transactionHandler *TransactionHandler,
) *MainHandler {
	return &MainHandler{
		LoginHandler: loginHandler,
		StaffHandler:  staffHandler,
		MenuHandler: menuHandler,
		ReservationHandler: reservationHandler,
		TransactionHandler: transactionHandler,
	}
}

func (h *MainHandler) Run() {

	staff, err := h.LoginHandler.Login()
    if err != nil {
        fmt.Println("❌ Login gagal:", err)
        return
    }

    staffID := staff.ID
    fmt.Printf("✅ Selamat datang, %s!\n", staff.Name)

	for {
		prompt := promptui.Select{
			Label: "🏪 Beverage CLI Main Menu",
			Items: []string{
				"1. Staff Management",
				"2. Menu Management",
				"3. Reservation Management",
				"4. Transaction Management",
				"5. Exit",
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
		case "4. Transaction Management":
			h.TransactionHandler.menu(staffID)
		case "5. Exit":
			fmt.Println("👋 Exiting Beverage CLI... Goodbye!")
			return
		}
	}
}
