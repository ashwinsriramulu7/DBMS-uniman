package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateCollege(newCollege models.College) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO college(name, location, estd) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Failed to prepare statement %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newCollege.Name, newCollege.Location, newCollege.Estd)
	if err != nil {
		log.Fatalf("Failed to exectute statement: %v", err)
	}
	fmt.Println("College added successfully")
}
func GetClassScheduleByID(id int) models.ClassSchedule {
	db := includes.InitDB()
	defer db.Close()
	var cs models.ClassSchedule
	err := db.QueryRow("SELECT * FROM class_schedule WHERE id = ?", id).Scan(
		&cs.ID, &cs.CourseID, &cs.FacultyID, &cs.DayOfWeek, &cs.StartTime, &cs.EndTime, &cs.Location)
	if err != nil {
		log.Fatal(err)
	}
	return cs
}

func DeleteClassScheduleByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM class_schedule WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

