package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateTeaches(t models.Teaches) {
	db := includes.InitDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO teaches(faculty_id, course_id, semester, year) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Prepare failed: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.FacultyID, t.CourseID, t.Semester, t.Year)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	fmt.Println("Teaches entry added.")
}
func GetTeachesByID(id int) models.Teaches {
	db := includes.InitDB()
	defer db.Close()
	var t models.Teaches
	err := db.QueryRow("SELECT * FROM teaches WHERE id = ?", id).Scan(
		&t.ID, &t.FacultyID, &t.CourseID, &t.Semester, &t.Year)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func DeleteTeachesByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM teaches WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

