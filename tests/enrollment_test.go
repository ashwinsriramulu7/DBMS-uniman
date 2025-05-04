package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateEnrollment(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	e := models.Enrollment{StudentID: 1, CourseID: 1, Semester: "Spring", Year: 2024, Grade: "A"}
	modules.CreateEnrollment(e)

	var grade string
	err := db.QueryRow("SELECT grade FROM enrollment WHERE student_id = 1 AND course_id = 1").Scan(&grade)
	if err != nil || grade != "A" {
		t.Error("Enrollment insert failed")
	}

	db.Exec("DELETE FROM enrollment WHERE student_id = 1 AND course_id = 1")
}
