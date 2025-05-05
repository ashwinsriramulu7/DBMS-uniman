package modules

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateEnrollment(e models.Enrollment) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO enrollment(student_id, course_id, semester, year, grade) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.StudentID, e.CourseID, e.Semester, e.Year, e.Grade)
	if err != nil {
		log.Fatal(err)
	}
}
func GetEnrollmentByID(id int) models.Enrollment {
	db := includes.InitDB()
	defer db.Close()
	var e models.Enrollment
	err := db.QueryRow("SELECT * FROM enrollment WHERE id = ?", id).Scan(
		&e.ID, &e.StudentID, &e.CourseID, &e.Semester, &e.Year, &e.Grade)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func DeleteEnrollmentByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM enrollment WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

