package tests
import (
	"github.com/ashwin/dbms-uniman/includes"
	"testing"
	"fmt"
)
func TestDatabaseConnection(t *testing.T) {
    db := includes.InitDB() 
    defer db.Close()    
    err := db.Ping()
    if err != nil {
        t.Fatalf("Failed to ping database: %v", err)
    }
}

