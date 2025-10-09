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

	// ==== STAFF ====
	staffRepo := repository.NewStaffRepository(conn)
	staffService := service.NewStaffService(staffRepo)
	staffHandler := handler.NewStaffHandler(staffService)

	// ==== MAIN HANDLER ====
	mainHandler := handler.NewMainHandler(staffHandler)
	mainHandler.Run()
}

