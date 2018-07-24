package loyalty

import (
	"fmt"
	"strings"
)

func SetUseQR(account string, qr string, point int) bool {

	db := ConnectDatabase()
	if db == nil {
		return false
	}
	defer db.Close()
	if strings.Split(qr, "|")[1] == "saving_account" {
		queryString := fmt.Sprintf("UPDATE account set SavingAccount=1 WHERE AccountID='%s'", account)
		rows, error := db.Query(queryString)
		if error != nil {
			return false
		}
		rows.Close()
	} else if strings.Split(qr, "|")[1] == "fixed_account" {
		queryString := fmt.Sprintf("UPDATE account set FixedAccount=1 WHERE AccountID='%s'", account)
		rows, error := db.Query(queryString)
		if error != nil {
			return false
		}
		rows.Close()
	}

	idqr := GetIDQR(qr)
	queryString := fmt.Sprintf("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('%s',%d,%d)", account, idqr, point)
	rows, error := db.Query(queryString)
	if error != nil {
		return false
	}
	rows.Close()

	return true
}