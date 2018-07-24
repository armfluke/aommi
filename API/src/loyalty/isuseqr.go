package loyalty

import (
	"database/sql"
	"fmt"
	"strings"
)

func QueryAndCheck(queryString string, db *sql.DB) int {
	rows, error := db.Query(queryString)
	if error != nil {
		return 0
	}
	if error := rows.Err(); error != nil {
		return 0
	}
	count := 0
	if rows.Next() {
		if error := rows.Scan(&count); error != nil {
			return 0
		}
	} else {
		count = 0
	}
	return count
}
func IsUseQR(account string, qr string) bool {
	db := ConnectDatabase()
	if db == nil {
		return false
	}
	defer db.Close()

	if strings.Split(qr, "|")[1] == "deposit" {
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM codeUsed WHERE AccountID='%s' AND DATE(DateEarn) = CURDATE()", account)

		countdeposit := QueryAndCheck(queryString, db)
		if countdeposit > 2 {
			return false
		}
	} else if strings.Split(qr, "|")[1] == "saving_account" {
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM account WHERE AccountID='%s' AND SavingAccount=0", account)

		count := QueryAndCheck(queryString, db)
		if count == 0 {
			return false
		}
	} else if strings.Split(qr, "|")[1] == "fixed_account" {
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM account WHERE AccountID='%s' AND FixedAccount=0", account)

		count := QueryAndCheck(queryString, db)
		if count == 0 {
			return false
		}
	}
	return true
}
