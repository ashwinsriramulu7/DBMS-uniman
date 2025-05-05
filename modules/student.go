package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
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
func GetStudentByID(id int) models.Student {
	db := includes.InitDB()
	defer db.Close()
	var s models.Student
	err := db.QueryRow("SELECT * FROM student WHERE id = ?", id).Scan(
		&s.ID, &s.Name, &s.MobileNumber, &s.Email, &s.ProgramEnrolled, &s.Type)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func DeleteStudentByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM student WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

