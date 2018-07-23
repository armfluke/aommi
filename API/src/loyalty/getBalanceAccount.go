package loyalty

import (
	"database/sql"
	"fmt"
)

func GetBalanceAccount(account string) (int,bool){
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if error != nil {
		return -1,false
	}

	defer db.Close()
	queryString := fmt.Sprintf("SELECT PointBalance FROM account WHERE AccountID='%s'",account)
	rows, error := db.Query(queryString)
	if err := rows.Err(); err != nil {
		return -1,false
	}
	balance := 0
	if rows.Next() {
		if err := rows.Scan(&balance); err != nil {
			return -1,false	
		}
		return balance,true
	}else{
		return balance,false
	}
	rows.Close()
	return balance,false	
}
