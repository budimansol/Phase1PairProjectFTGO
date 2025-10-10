package main

import (
	"github.com/budimansol/pairproject/db"
	"github.com/budimansol/pairproject/internal/handler"
	"github.com/budimansol/pairproject/internal/repository"
	"github.com/budimansol/pairproject/internal/service"
)

func main() {
	conn := db.ConnectDB()
	defer conn.Close()

	// --- Repositories ---
	staffRepo := repository.NewStaffRepository(conn)
	menuRepo := repository.NewMenuRepository(conn)
	memberRepo := repository.NewMemberRepository(conn)
	resRepo := repository.NewReservationRepository(conn)
	transRepo := repository.NewTransactionRepository(conn)

	// --- Services ---
	staffService := service.NewStaffService(staffRepo)
	menuService := service.NewMenuService(menuRepo)
	memberService := service.NewMemberService(memberRepo)
	resService := service.NewReservationService(resRepo)
	transService := service.NewTransactionService(transRepo)

	// --- Login Handler ---
	loginHandler := handler.NewLoginHandler(staffService)

	// --- Handlers ---
	staffHandler := handler.NewStaffHandler(staffService)
	menuHandler := handler.NewMenuHandler(menuService)
	resHandler := handler.NewReservationHandler(resService, memberService)
	transHandler := handler.NewTransactionHandler(transService, memberService, menuService, staffService)

	// --- Main Handler ---
	mainHandler := handler.NewMainHandler(
		loginHandler,
		staffHandler,
		menuHandler,
		resHandler,
		transHandler,
	)

	// --- Run CLI ---
	mainHandler.Run()
}
