package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
)

func TestCreateCollege(t *testing.T) {
	db := includes.InitDB()
	defer db.Close()
	testCollege := models.College{
		Name:     "Test College",
		Location: "Test Location",
		Estd:     2020,
	}
	modules.CreateCollege(testCollege)
	var name, location string
	var estd int
	err := db.QueryRow("SELECT name, location, estd FROM college WHERE name = ?", "Test College").Scan(
		&name, &location, &estd)
	if err != nil {
		t.Fatalf("Failed to fetch inserted college: %v", err)
	}
	if name != testCollege.Name || location != testCollege.Location || estd != testCollege.Estd {
		t.Errorf("Inserted data does not match.\nExpected: %+v\nGot: name=%s, location=%s, estd=%d", testCollege, name, location, estd)
	}
	db.Exec("DELETE FROM college WHERE name = 'Test College'")
}
func TestGetAndDeleteCollege(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    testCollege := models.College{Name: "GetDelete College", Location: "Nowhere", Estd: 2021}
    modules.CreateCollege(testCollege)

    var id int
    err := db.QueryRow("SELECT id FROM college WHERE name = ?", testCollege.Name).Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted college ID: %v", err)
    }

    // Implement GetCollegeByID in your modules if not already present
    got := modules.GetCollegeByID(id)
    if got.Name != testCollege.Name || got.Location != testCollege.Location || got.Estd != testCollege.Estd {
        t.Errorf("GetCollegeByID failed: expected %+v, got %+v", testCollege, got)
    }

    modules.DeleteCollegeByID(id)
    err = db.QueryRow("SELECT id FROM college WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteCollegeByID failed: record still exists")
    }
}

