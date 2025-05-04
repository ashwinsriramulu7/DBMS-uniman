package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateDepartment(d models.Department) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO department(name, head_of_department, college) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Prepare failed: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.Name, d.HeadOfDepartment, d.College)
	if err != nil {
		log.Fatalf("Exec failed: %v", err)
	}
	fmt.Println("Department added successfully")
}
