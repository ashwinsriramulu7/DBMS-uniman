package tests

import (
	"testing"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
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

