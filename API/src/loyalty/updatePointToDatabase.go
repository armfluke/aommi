package loyalty

import (
	"database/sql"
	"fmt"
)

func updatePointToDatabase(balance int,account string) (bool) {
	
	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		return false
	}

	defer db.Close()

	queryString := fmt.Sprintf("UPDATE account SET PointBalance = %d WHERE AccountID = '%s'",balance,account)
	results, err := db.Query(queryString)
	results.Close()

	if err != nil {
		return false
	}
	return true
}