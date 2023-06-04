package main

import (
	"fmt"
	"os"

	"gorm-tutorial/database"
)

func main() {

	db, err := database.GetGormConn()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	// Close database connection
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sqlDB.Close()
	fmt.Println("Database connection closed")

}
