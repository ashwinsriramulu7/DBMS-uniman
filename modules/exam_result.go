package modules

import (
    "log"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func CreateExamResult(er models.ExamResult) {
    db := includes.InitDB()
    defer db.Close()
    stmt, err := db.Prepare("INSERT INTO exam_result(exam_id, student_id, marks_obtained) VALUES (?, ?, ?)")
    if err != nil { log.Fatal(err) }
    defer stmt.Close()
    _, err = stmt.Exec(er.ExamID, er.StudentID, er.MarksObtained)
    if err != nil { log.Fatal(err) }
}

