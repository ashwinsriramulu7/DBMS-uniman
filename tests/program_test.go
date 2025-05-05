package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
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
func TestGetAndDeleteProgram(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    prog := models.Program{Name: "GetDelete Program", Level: "PG", DepartmentID: 1}
    modules.CreateProgram(prog)

    var id int
    err := db.QueryRow("SELECT id FROM program WHERE name = ?", prog.Name).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted program ID: %v", err)
    }

    got := modules.GetProgramByID(id)
    if got.Name != prog.Name || got.Level != prog.Level || got.DepartmentID != prog.DepartmentID {
        t.Errorf("GetProgramByID failed: expected %+v, got %+v", prog, got)
    }

    modules.DeleteProgramByID(id)
    err = db.QueryRow("SELECT id FROM program WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteProgramByID failed: record still exists")
    }
}

