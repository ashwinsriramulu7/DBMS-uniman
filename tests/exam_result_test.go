package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateExamResult(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	er := models.ExamResult{ExamID: 14, StudentID: 1, MarksObtained: 85}
	modules.CreateExamResult(er)

	var marks int
	err := db.QueryRow("SELECT marks_obtained FROM exam_result WHERE exam_id = 1 AND student_id = 1").Scan(&marks)
	if err != nil || marks != 85 {
		t.Error("ExamResult insert failed")
	}

	db.Exec("DELETE FROM exam_result WHERE exam_id = 1 AND student_id = 1")
}
