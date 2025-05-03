package modules
import (
	"fmt"
	"log"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
)
func CreateFaculty(f models.Faculty) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO faculty(name, mobile_number, email, address, type, title) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(f.Name, f.MobileNumber, f.Email, f.Address, f.Type, f.Title)
	if err != nil {
		log.Fatalf("Failed to execute statement: %v", err)
	}
	fmt.Println("Faculty added successfully")
}

