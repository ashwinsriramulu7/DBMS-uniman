package modules

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateProgram(p models.Program) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO program(name, level, department_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Level, p.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}
}
func GetProgramByID(id int) models.Program {
	db := includes.InitDB()
	defer db.Close()
	var p models.Program
	err := db.QueryRow("SELECT * FROM program WHERE id = ?", id).Scan(
		&p.ID, &p.Name, &p.Level, &p.DepartmentID)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func DeleteProgramByID(id int) {
	db := includes.InitDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM program WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

