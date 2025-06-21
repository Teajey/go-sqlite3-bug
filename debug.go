package bug

import (
	"database/sql"
	"fmt"
	"log"
)

func ListTables(d *sql.DB) []string {
	rows, err := d.Query("SELECT name FROM sqlite_master WHERE type = 'table';")
	if err != nil {
		log.Println("Query failed: %w", err)
		return nil
	}
	defer rows.Close()
	names := []string{}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan failed: %w", err)
			return nil
		}
		names = append(names, name)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Rows returned error: %w", err)
		return nil
	}
	return names
}
