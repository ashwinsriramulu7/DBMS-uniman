package tests

import (
	"testing"
	"log"
	"github.com/ashwin/dbms-uniman/includes"
)

func TestCleanupColleges(t *testing.T) {
	// Insert mock test data manually (setup)
	db := includes.InitDB()
	defer db.Close()

	// Insert some test records
	_, err := db.Exec(`INSERT INTO college (name) VALUES ('Test College X'), ('Another Test College Y')`)
	if err != nil {
		t.Fatalf("Failed to insert test records: %v", err)
	}

	// Run the cleanup
	err = includes.CleanupColleges()
	if err != nil {
		t.Fatalf("CleanupColleges failed: %v", err)
	}

	// Check that the test data is gone
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM college WHERE name LIKE '%Test%'`).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count remaining test records: %v", err)
	}

	if count != 0 {
		t.Errorf("Expected 0 test colleges remaining, but got %d", count)
	} else {
		log.Println("TestCleanupColleges: All test records successfully deleted.")
	}
}

