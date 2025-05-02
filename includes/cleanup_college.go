package includes

import (
    "fmt"
    "log"
)
func CleanupColleges() error {
	db:= InitDB()
    defer db.Close()

    query := `DELETE FROM college WHERE name LIKE ?`
    res, err := db.Exec(query, "%Test%")
    if err != nil {
        return fmt.Errorf("failed to clean up test colleges: %v", err)
    }

    rowsAffected, _ := res.RowsAffected()
    log.Printf("CleanupColleges: %d test records deleted.\n", rowsAffected)
    return nil
}

