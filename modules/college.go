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
func GetCollegeByID(id int) models.College {
	db := includes.InitDB()
	defer db.Close()
	var c models.College
	err := db.QueryRow("SELECT * FROM college WHERE id = ?", id).Scan(&c.ID, &c.Name, &c.Location, &c.Estd)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func DeleteCollegeByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM college WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

