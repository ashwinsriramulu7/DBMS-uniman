package modules

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateExam(e models.Exam) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO exam(course_id, exam_type, date, total_marks) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.CourseID, e.ExamType, e.Date, e.TotalMarks)
	if err != nil {
		log.Fatal(err)
	}
}
func GetExamByID(id int) models.Exam {
	db := includes.InitDB()
	defer db.Close()
	var e models.Exam
	err := db.QueryRow("SELECT * FROM exam WHERE id = ?", id).Scan(
		&e.ID, &e.CourseID, &e.ExamType, &e.Date, &e.TotalMarks)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func DeleteExamByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM exam WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

