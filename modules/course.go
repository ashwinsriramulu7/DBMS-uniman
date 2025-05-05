package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateCourse(c models.Course) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO course(department_id, course_code) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("Prepare failed: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.DepartmentID, c.CourseCode)
	if err != nil {
		log.Fatalf("Exec failed: %v", err)
	}
	fmt.Println("Course added successfully")
}
func GetCourseByID(id int) models.Course {
	db := includes.InitDB()
	defer db.Close()
	var c models.Course
	err := db.QueryRow("SELECT * FROM course WHERE id = ?", id).Scan(&c.ID, &c.DepartmentID, &c.CourseCode)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func DeleteCourseByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM course WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

