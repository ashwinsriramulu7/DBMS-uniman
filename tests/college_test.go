package tests

import (
	"testing"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
)
func TestCreateCollege(t *testing.T){
	db := includes.InitDB()
	defer db.Close()
	testCollege := models.College{
		Name : "Test College",
		Location : "Test Location",
		Estd : 2020,
	}
	modules.CreateCollege(testCollege)
	var name, location string
	var estd int
	err := db.QueryRow("SELECT name, location, estd FROM college WHERE name = ?","Test College").Scan(
		&name, &location, &estd)
	if err != nil{
		t.Fatalf("Failed to fetch inserted college: %v", err)
	}
	if name != testCollege.Name || location != testCollege.Location || estd != testCollege.Estd{
        	t.Errorf("Inserted data does not match.\nExpected: %+v\nGot: name=%s, location=%s, estd=%d", testCollege, name, location, estd)
    	}
	db.Exec("DELETE FROM college WHERE name = 'Test College'")
}


