package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateFaculty(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()

	testFaculty := models.Faculty{
		Name:         "Test Faculty",
		MobileNumber: "9999999999",
		Email:        "testfaculty@example.com",
		Address:      "Test Address",
		Type:         "ACADEMIC",
		Title:        "Professor",
	}
	modules.CreateFaculty(testFaculty)

	var name, mobile, email, address, ftype, title string
	err := db.QueryRow("SELECT name, mobile_number, email, address, type, title FROM faculty WHERE email = ?", testFaculty.Email).
		Scan(&name, &mobile, &email, &address, &ftype, &title)
	if err != nil {
		t.Fatalf("Failed to fetch faculty: %v", err)
	}
	if name != testFaculty.Name || mobile != testFaculty.MobileNumber || email != testFaculty.Email || address != testFaculty.Address || ftype != testFaculty.Type || title != testFaculty.Title {
		t.Errorf("Mismatch in faculty data")
	}
	db.Exec("DELETE FROM faculty WHERE email = ?", testFaculty.Email)
}
func TestGetAndDeleteFaculty(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    testFaculty := models.Faculty{
        Name: "GetDelete Faculty", MobileNumber: "7777777777", Email: "getdeletefaculty@example.com",
        Address: "Somewhere", Type: "ACADEMIC", Title: "Assistant Professor",
    }
    modules.CreateFaculty(testFaculty)

    var id int
    err := db.QueryRow("SELECT id FROM faculty WHERE email = ?", testFaculty.Email).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted faculty ID: %v", err)
    }

    got := modules.GetFacultyByID(id)
    if got.Name != testFaculty.Name || got.MobileNumber != testFaculty.MobileNumber || got.Email != testFaculty.Email ||
        got.Address != testFaculty.Address || got.Type != testFaculty.Type || got.Title != testFaculty.Title {
        t.Errorf("GetFacultyByID failed: expected %+v, got %+v", testFaculty, got)
    }

    modules.DeleteFacultyByID(id)
    err = db.QueryRow("SELECT id FROM faculty WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteFacultyByID failed: record still exists")
    }
}

