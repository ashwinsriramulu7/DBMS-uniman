package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
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
func GetFacultyByID(id int) models.Faculty {
	db := includes.InitDB()
	defer db.Close()
	var f models.Faculty
	err := db.QueryRow("SELECT * FROM faculty WHERE id = ?", id).Scan(
		&f.ID, &f.Name, &f.MobileNumber, &f.Email, &f.Address, &f.Type, &f.Title)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func DeleteFacultyByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM faculty WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

