package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
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
func TestGetAndDeleteClassSchedule(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    // Insert
    cs := models.ClassSchedule{CourseID: 1, FacultyID: 1, DayOfWeek: "Tuesday", StartTime: "12:00:00", EndTime: "13:00:00", Location: "Room 202"}
    modules.CreateClassSchedule(cs)

    // Get ID
    var id int
    err := db.QueryRow("SELECT id FROM class_schedule WHERE course_id = 1 AND faculty_id = 1 AND location = 'Room 202'").Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted class_schedule ID: %v", err)
    }

    // Get
    got := modules.GetClassScheduleByID(id)
    if got.Location != cs.Location || got.CourseID != cs.CourseID || got.FacultyID != cs.FacultyID {
        t.Errorf("GetClassScheduleByID failed: expected %+v, got %+v", cs, got)
    }

    // Delete
    modules.DeleteClassScheduleByID(id)
    err = db.QueryRow("SELECT id FROM class_schedule WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteClassScheduleByID failed: record still exists")
    }
}

