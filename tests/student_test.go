package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	testStudent := models.Student{
		Name:            "Test Student",
		MobileNumber:    "8888888888",
		Email:           "teststudent@example.com",
		ProgramEnrolled: "B.Tech",
		Type:            "UG",
	}
	modules.CreateStudent(testStudent)

	var name, mobile, email, program, stype string
	err := db.QueryRow("SELECT name, mobile_number, email, program_enrolled, type FROM student WHERE email = ?", testStudent.Email).
		Scan(&name, &mobile, &email, &program, &stype)
	if err != nil {
		t.Fatalf("Failed to fetch student: %v", err)
	}
	if name != testStudent.Name || mobile != testStudent.MobileNumber || email != testStudent.Email || program != testStudent.ProgramEnrolled || stype != testStudent.Type {
		t.Errorf("Mismatch in student data")
	}
	db.Exec("DELETE FROM student WHERE email = ?", testStudent.Email)
}
func TestGetAndDeleteStudent(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    testStudent := models.Student{
        Name: "GetDelete Student", MobileNumber: "6666666666", Email: "getdeletestudent@example.com",
        ProgramEnrolled: "M.Tech", Type: "PG",
    }
    modules.CreateStudent(testStudent)

    var id int
    err := db.QueryRow("SELECT id FROM student WHERE email = ?", testStudent.Email).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted student ID: %v", err)
    }

    got := modules.GetStudentByID(id)
    if got.Name != testStudent.Name || got.MobileNumber != testStudent.MobileNumber || got.Email != testStudent.Email ||
        got.ProgramEnrolled != testStudent.ProgramEnrolled || got.Type != testStudent.Type {
        t.Errorf("GetStudentByID failed: expected %+v, got %+v", testStudent, got)
    }

    modules.DeleteStudentByID(id)
    err = db.QueryRow("SELECT id FROM student WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteStudentByID failed: record still exists")
    }
}

