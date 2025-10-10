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

	// Login Handler
	staffRepo := repository.NewStaffRepository(conn)
	staffService := service.NewStaffService(staffRepo)
	// loginHandler := handler.NewLoginHandler(staffService)
	staffHandler := handler.NewStaffHandler(staffService)

	// staffName, err := loginHandler.Login()
	// if err != nil {
	// 	fmt.Println("Login error:", err)
	// 	return
	// }

	// fmt.Printf("🎉 Logged in as: %s\n\n", *staffName)

	// Setelah login, load handler lain

	// Repository & Service

	menuRepo := repository.NewMenuRepository(conn)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuService)

	memberRepo := repository.NewMemberRepository(conn)
	memberService := service.NewMemberService(memberRepo)
	memberHandler := handler.NewMemberHandler(memberService)

	resRepo := repository.NewReservationRepository(conn)
	resService := service.NewReservationService(resRepo)
	resHandler := handler.NewReservationHandler(resService, memberService)

	// ==== REPORTS ====
	reportRepo := repository.NewReportRepository(conn)
	reportService := service.NewReportService(reportRepo)
	reportHandler := handler.NewReportHandler(reportService)

	// ==== MAIN HANDLER ====
	mainHandler := handler.NewMainHandler(staffHandler, menuHandler, resHandler, memberHandler, reportHandler)
	mainHandler.Run()
}
