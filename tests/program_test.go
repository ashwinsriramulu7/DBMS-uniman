package tests

import (
    "testing"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/modules"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func TestCreateProgram(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    prog := models.Program{Name: "B.Tech CSE", Level: "UG", DepartmentID: 1}
    modules.CreateProgram(prog)

    var name string
    err := db.QueryRow("SELECT name FROM program WHERE name = ?", "B.Tech CSE").Scan(&name)
    if err != nil || name != "B.Tech CSE" {
        t.Error("Program insert failed")
    }

    db.Exec("DELETE FROM program WHERE name = 'B.Tech CSE'")
}

