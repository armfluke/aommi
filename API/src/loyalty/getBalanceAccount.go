package loyalty

import (
	"fmt"
)

func GetBalanceAccount(account string) (int, bool) {
	db := ConnectDatabase()
	if db == nil {
		return -1, false
	}

	defer db.Close()
	queryString := fmt.Sprintf("SELECT PointBalance FROM account WHERE AccountID='%s'", account)
	rows, _ := db.Query(queryString)
	if err := rows.Err(); err != nil {
		return -1, false
	}
	balance := 0
	if rows.Next() {
		if err := rows.Scan(&balance); err != nil {
			return -1, false
		}
		return balance, true
	} else {
		return balance, false
	}
	rows.Close()
	return balance, false
}
