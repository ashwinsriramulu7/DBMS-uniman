package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
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
func TestGetAndDeleteCourse(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    testCourse := models.Course{DepartmentID: 1, CourseCode: "CS102"}
    modules.CreateCourse(testCourse)

    var id int
    err := db.QueryRow("SELECT id FROM course WHERE course_code = ?", testCourse.CourseCode).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted course ID: %v", err)
    }

    got := modules.GetCourseByID(id)
    if got.DepartmentID != testCourse.DepartmentID || got.CourseCode != testCourse.CourseCode {
        t.Errorf("GetCourseByID failed: expected %+v, got %+v", testCourse, got)
    }

    modules.DeleteCourseByID(id)
    err = db.QueryRow("SELECT id FROM course WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteCourseByID failed: record still exists")
    }
}

