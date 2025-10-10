package main

import (
	"fmt"

	"github.com/budimansol/pairproject/db"
	"github.com/budimansol/pairproject/internal/handler"
	"github.com/budimansol/pairproject/internal/repository"
	"github.com/budimansol/pairproject/internal/service"
)

func main() {
	conn := db.ConnectDB()
	defer conn.Close()

	// Repository & Service
	staffRepo := repository.NewStaffRepository(conn)
	staffService := service.NewStaffService(staffRepo)

	// Login Handler
	loginHandler := handler.NewLoginHandler(staffService)
	staffName, err := loginHandler.Login()
	if err != nil {
		fmt.Println("Login error:", err)
		return
	}

	fmt.Printf("🎉 Logged in as: %s\n\n", *staffName)

	// Setelah login, load handler lain
	staffHandler := handler.NewStaffHandler(staffService)

	menuRepo := repository.NewMenuRepository(conn)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuService)


	mainHandler := handler.NewMainHandler(staffHandler, menuHandler)
	mainHandler.Run()
}


