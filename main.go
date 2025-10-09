package main

import (
	"fmt"
	"log"

	"github.com/budimansol/pairproject/config"
)

func main() {
	config.ConnectDB()

	rows, err := config.DB.Query("SELECT * FROM staffs")
	if err != nil {
		log.Fatal("Query error:", err)
	}
	defer rows.Close()

	fmt.Println("📋 Staff List:")
	for rows.Next() {
		var id int
		var name, email, password, role string
		var createdAt string
		rows.Scan(&id, &name, &email, &password, &role, &createdAt)
		fmt.Printf("%d | %s | %s | %s | %s\n", id, name, email, password, role)
	}
}
