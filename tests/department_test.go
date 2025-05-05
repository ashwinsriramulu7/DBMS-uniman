package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateDepartment(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	testDepartment := models.Department{
		Name:             "Test Department",
		HeadOfDepartment: 1, // Assumes a faculty with id=1 exists
		College:          1, // Assumes a college with id=1 exists
	}
	modules.CreateDepartment(testDepartment)

	var name string
	var hod, college int
	err := db.QueryRow("SELECT name, head_of_department, college FROM department WHERE name = ?", testDepartment.Name).
		Scan(&name, &hod, &college)
	if err != nil {
		t.Fatalf("Failed to fetch department: %v", err)
	}
	if name != testDepartment.Name || hod != testDepartment.HeadOfDepartment || college != testDepartment.College {
		t.Errorf("Mismatch in department data")
	}
	db.Exec("DELETE FROM department WHERE name = ?", testDepartment.Name)
}
func TestGetAndDeleteDepartment(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    testDepartment := models.Department{Name: "GetDelete Dept", HeadOfDepartment: 1, College: 1}
    modules.CreateDepartment(testDepartment)

    var id int
    err := db.QueryRow("SELECT id FROM department WHERE name = ?", testDepartment.Name).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted department ID: %v", err)
    }

    got := modules.GetDepartmentByID(id)
    if got.Name != testDepartment.Name || got.HeadOfDepartment != testDepartment.HeadOfDepartment || got.College != testDepartment.College {
        t.Errorf("GetDepartmentByID failed: expected %+v, got %+v", testDepartment, got)
    }

    modules.DeleteDepartmentByID(id)
    err = db.QueryRow("SELECT id FROM department WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteDepartmentByID failed: record still exists")
    }
}

