package modules

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateExamResult(er models.ExamResult) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO exam_result(exam_id, student_id, marks_obtained) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(er.ExamID, er.StudentID, er.MarksObtained)
	if err != nil {
		log.Fatal(err)
	}
}
func GetExamResultByID(id int) models.ExamResult {
	db := includes.InitDB()
	defer db.Close()
	var er models.ExamResult
	err := db.QueryRow("SELECT * FROM exam_result WHERE id = ?", id).Scan(
		&er.ID, &er.ExamID, &er.StudentID, &er.MarksObtained)
	if err != nil {
		log.Fatal(err)
	}
	return er
}

func DeleteExamResultByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM exam_result WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

