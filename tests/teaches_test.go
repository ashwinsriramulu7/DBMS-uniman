package tests

import (
    "testing"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/modules"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func TestCreateTeaches(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    entry := models.Teaches{FacultyID: 1, CourseID: 1, Semester: "Fall", Year: 2023}
    modules.CreateTeaches(entry)

    var semester string
    err := db.QueryRow("SELECT semester FROM teaches WHERE faculty_id = ? AND course_id = ?", 1, 1).Scan(&semester)
    if err != nil || semester != "Fall" {
        t.Errorf("Teaches insert failed")
    }

    db.Exec("DELETE FROM teaches WHERE faculty_id = 1 AND course_id = 1")
}

