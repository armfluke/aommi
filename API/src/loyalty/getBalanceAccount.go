package loyalty

import (
	"database/sql"
)

func GetBalanceAccount(account string) (int,bool){
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if error != nil {
		return -1,false
	}

	defer db.Close()

	rows, error := db.Query("SELECT PointBalance FROM account")
	if err := rows.Err(); err != nil {
		return -1,false
	}
	var balance int
	for rows.Next() {
		if err := rows.Scan(&balance); err != nil {
			return -1,false	
		}
	}
	rows.Close()
	
	return balance,true
}
