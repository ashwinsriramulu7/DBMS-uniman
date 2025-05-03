package modules
import(
	"fmt"
	"log"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
)
func CreateCollege(newCollege models.College){
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO college(name, location, estd) VALUES (?, ?, ?)")
	if err != nil{
		log.Fatalf("Failed to prepare statement %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newCollege.Name, newCollege.Location, newCollege.Estd)
	if err != nil{
		log.Fatalf("Failed to exectute statement: %v", err)
	}
	fmt.Println("College added successfully")
}
	
