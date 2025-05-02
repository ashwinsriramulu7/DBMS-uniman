package modules

import (
	"fmt"
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
)

func AddCollege(data map[string]any) error {
	db := includes.InitDB()
	defer db.Close()

	stmt := `INSERT INTO college(name, location, estd) VALUES (?, ?, ?)`
	_, err := db.Exec(stmt, data["name"], data["location"], data["estd"])
	if err != nil {
		fmt.Println("AddCollege error:", err)
		return err
	}
	return nil
}

