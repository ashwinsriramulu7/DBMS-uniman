package tests

import (
    "testing"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/modules"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func TestCreateExam(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    e := models.Exam{CourseID: 1, ExamType: "MIDTERM", Date: "2024-05-01", TotalMarks: 100}
    modules.CreateExam(e)

    var marks int
    err := db.QueryRow("SELECT total_marks FROM exam WHERE course_id = 1 AND exam_type = 'MIDTERM'").Scan(&marks)
    if err != nil || marks != 100 {
        t.Error("Exam insert failed")
    }

    db.Exec("DELETE FROM exam WHERE course_id = 1 AND exam_type = 'MIDTERM'")
}

