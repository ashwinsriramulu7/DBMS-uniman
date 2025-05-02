package includes
import(
	"fmt" 
	"database/sql"
	"db.go"
)
func cleanupColleges() error{
	db = includes.ConnectDB()
	drop from college where name contains "Test";
}
