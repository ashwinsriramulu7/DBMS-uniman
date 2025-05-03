package modules

import (
	"fmt"
	"log"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
)

func CreateStudent(s models.Student) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO student(name, mobile_number, email, program_enrolled, type) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Prepare error: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Name, s.MobileNumber, s.Email, s.ProgramEnrolled, s.Type)
	if err != nil {
		log.Fatalf("Exec error: %v", err)
	}
	fmt.Println("Student added successfully")
}

