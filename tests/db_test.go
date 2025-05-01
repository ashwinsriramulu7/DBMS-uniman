package tests
import (
	"/mnt/c/Users/Ashwin/Documents/DBMS-uniman/includes"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestDBConnection(t *testing.T){
	db := includes.InitDB
	defer db.close()
	err := db.Ping()
	assert.NoError(t, err, "Expected no error when pinging the database")
}
