package tests
import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"testing"
)
func TestDatabaseConnection(t *testing.T) {
    db := includes.InitDB() 
    defer db.Close()    
    err := db.Ping()
    if err != nil {
        t.Fatalf("Failed to ping database: %v", err)
    }
}

