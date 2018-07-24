package loyalty

import (
	"database/sql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")
	if err != nil {
		return nil
	}
	
	return db
}