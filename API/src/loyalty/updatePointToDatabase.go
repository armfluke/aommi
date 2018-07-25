package loyalty

import (
	"fmt"
)

func UpdatePointToDatabase(balance int, account string) bool {

	db := ConnectDatabase()

	if db == nil {
		return false
	}
	defer db.Close()

	queryString := fmt.Sprintf("UPDATE account SET PointBalance = %d WHERE AccountID = '%s'", balance, account)
	results, error := db.Query(queryString)
	results.Close()

	if error != nil {
		return false
	}
	return true
}
