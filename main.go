package main

import (
	"database/sql"
	"fmt"
	"log" // Added for cleaner error handling exit

	_ "github.com/mattn/go-sqlite3"

	"fieldConnector/returnVisit"
)

func main() {
	fmt.Println("\n WELCOME! TO FIELDCONNECTOR APP") 
	fmt.Println()
	fmt.Println("\nOPTIONS")

	db, err := sql.Open("sqlite3", "return_visit.db")
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("1. Add Return Visit\n2. Show All Return Visit\n3. Delete From Return Visit\n4. Update Return Visit\n5. Filter By Name Or Street")

	var option string
	fmt.Println("Choose Your Option Please: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		returnVisit.UserInput(db)
	case "2":
		returnVisit.Display(db)
	case "3":
		returnVisit.Display(db)
		var id string
		fmt.Println("\nPlease, enter ID number you want to delete: ")
		fmt.Scanln(&id)
		returnVisit.DeleteUser(db, id)
	case "4":
		returnVisit.Display(db)
		fmt.Println()
		returnVisit.UpdateUser(db)
	case "5":
		returnVisit.FilterByFields(db)
	default:
		fmt.Println("Invalid option selected.") 
	}
}
