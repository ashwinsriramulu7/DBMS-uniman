package modules

import (
    "log"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func CreateClassSchedule(cs models.ClassSchedule) {
    db := includes.InitDB()
    defer db.Close()
    stmt, err := db.Prepare("INSERT INTO class_schedule(course_id, faculty_id, day_of_week, start_time, end_time, location) VALUES (?, ?, ?, ?, ?, ?)")
    if err != nil { log.Fatal(err) }
    defer stmt.Close()
    _, err = stmt.Exec(cs.CourseID, cs.FacultyID, cs.DayOfWeek, cs.StartTime, cs.EndTime, cs.Location)
    if err != nil { log.Fatal(err) }
}

