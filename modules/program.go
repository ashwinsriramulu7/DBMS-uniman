package modules

import (
    "log"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func CreateProgram(p models.Program) {
    db := includes.InitDB()
    defer db.Close()
    stmt, err := db.Prepare("INSERT INTO program(name, level, department_id) VALUES (?, ?, ?)")
    if err != nil { log.Fatal(err) }
    defer stmt.Close()
    _, err = stmt.Exec(p.Name, p.Level, p.DepartmentID)
    if err != nil { log.Fatal(err) }
}

