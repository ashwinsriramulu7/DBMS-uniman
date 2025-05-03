package tests

import (
    "testing"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/modules"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func TestCreateClassSchedule(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    cs := models.ClassSchedule{CourseID: 1, FacultyID: 1, DayOfWeek: "Monday", StartTime: "10:00:00", EndTime: "11:00:00", Location: "Room 101"}
    modules.CreateClassSchedule(cs)

    var loc string
    err := db.QueryRow("SELECT location FROM class_schedule WHERE course_id = 1 AND faculty_id = 1").Scan(&loc)
    if err != nil || loc != "Room 101" {
        t.Error("ClassSchedule insert failed")
    }

    db.Exec("DELETE FROM class_schedule WHERE course_id = 1 AND faculty_id = 1")
}

