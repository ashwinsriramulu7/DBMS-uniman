package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
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
func TestGetAndDeleteExam(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    e := models.Exam{CourseID: 1, ExamType: "FINAL", Date: "2025-06-01", TotalMarks: 150}
    modules.CreateExam(e)

    var id int
    err := db.QueryRow("SELECT id FROM exam WHERE course_id = 1 AND exam_type = 'FINAL'").Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted exam ID: %v", err)
    }

    got := modules.GetExamByID(id)
    if got.CourseID != e.CourseID || got.ExamType != e.ExamType || got.Date != e.Date || got.TotalMarks != e.TotalMarks {
        t.Errorf("GetExamByID failed: expected %+v, got %+v", e, got)
    }

    modules.DeleteExamByID(id)
    err = db.QueryRow("SELECT id FROM exam WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteExamByID failed: record still exists")
    }
}

