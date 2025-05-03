package modules

import (
    "log"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func CreateEnrollment(e models.Enrollment) {
    db := includes.InitDB()
    defer db.Close()
    stmt, err := db.Prepare("INSERT INTO enrollment(student_id, course_id, semester, year, grade) VALUES (?, ?, ?, ?, ?)")
    if err != nil { log.Fatal(err) }
    defer stmt.Close()
    _, err = stmt.Exec(e.StudentID, e.CourseID, e.Semester, e.Year, e.Grade)
    if err != nil { log.Fatal(err) }
}

