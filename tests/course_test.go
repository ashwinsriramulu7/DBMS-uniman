package tests

import (
	"testing"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func TestCreateCourse(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	testCourse := models.Course{
		DepartmentID: 1, // Assumes department with id=1 exists
		CourseCode:   "CS101",
	}
	modules.CreateCourse(testCourse)

	var deptID int
	var code string
	err := db.QueryRow("SELECT department_id, course_code FROM course WHERE course_code = ?", testCourse.CourseCode).
		Scan(&deptID, &code)
	if err != nil {
		t.Fatalf("Failed to fetch course: %v", err)
	}
	if deptID != testCourse.DepartmentID || code != testCourse.CourseCode {
		t.Errorf("Mismatch in course data")
	}
	db.Exec("DELETE FROM course WHERE course_code = ?", testCourse.CourseCode)
}

