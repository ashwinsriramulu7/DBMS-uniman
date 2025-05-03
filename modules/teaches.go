package modules
import (
    "log"
    "fmt"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func CreateTeaches(t models.Teaches) {
    db := includes.InitDB()
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO teaches(faculty_id, course_id, semester, year) VALUES (?, ?, ?, ?)")
    if err != nil {
        log.Fatalf("Prepare failed: %v", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(t.FacultyID, t.CourseID, t.Semester, t.Year)
    if err != nil {
        log.Fatalf("Insert failed: %v", err)
    }
    fmt.Println("Teaches entry added.")
}

