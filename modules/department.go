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
func GetDepartmentByID(id int) models.Department {
	db := includes.InitDB()
	defer db.Close()
	var d models.Department
	err := db.QueryRow("SELECT * FROM department WHERE id = ?", id).Scan(
		&d.ID, &d.Name, &d.HeadOfDepartment, &d.College)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func DeleteDepartmentByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM department WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

