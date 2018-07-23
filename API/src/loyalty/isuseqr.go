package loyalty

import (
	"database/sql"
	"fmt"
	"strings"
)

func IsUseQR(account string,qr string) bool{
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")
	if error != nil {
		return false
	}
	defer db.Close()

	var countdeposit int
	if strings.Split(qr,"|")[1]=="deposit"{
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM codeUsed WHERE AccountID='%s' AND DateEarn = CURDATE()",account)
		rows, _ := db.Query(queryString)
		if error := rows.Err(); error != nil {
			return false
		}
		if rows.Next() {
			if error := rows.Scan(&countdeposit); error != nil {
				return false	
			}
		}else{
			countdeposit = 0
		}
		if countdeposit>3 {
			return false
		}
	}else if strings.Split(qr,"|")[1]=="saving_account"{
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM account WHERE AccountID='%s' AND SavingAccount=0",account)
		rows, _ := db.Query(queryString)
		if error := rows.Err(); error != nil {
			return false
		}
		if rows.Next() {
			if error := rows.Scan(&countdeposit); error != nil {
				return false	
			}
		}else{
			countdeposit = 0
		}
		if countdeposit==0 {
			return false
		}
	}else if strings.Split(qr,"|")[1]=="fixed_account"{
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM account WHERE AccountID='%s' AND FixedAccount=0",account)
		rows, _ := db.Query(queryString)
		if error := rows.Err(); error != nil {
			return false
		}
		if rows.Next() {
			if error := rows.Scan(&countdeposit); error != nil {
				return false	
			}
		}else{
			countdeposit = 0
		}
		if countdeposit==0 {
			return false
		}
	}
	return true
}
