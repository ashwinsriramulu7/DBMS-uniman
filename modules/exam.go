package modules
import (
    "log"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func CreateExam(e models.Exam) {
    db := includes.InitDB()
    defer db.Close()
    stmt, err := db.Prepare("INSERT INTO exam(course_id, exam_type, date, total_marks) VALUES (?, ?, ?, ?)")
    if err != nil { log.Fatal(err) }
    defer stmt.Close()
    _, err = stmt.Exec(e.CourseID, e.ExamType, e.Date, e.TotalMarks)
    if err != nil { log.Fatal(err) }
}

