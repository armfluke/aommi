package loyalty

import (
	"database/sql"
	"fmt"
)

func GetIDQR(qr string) int{
	
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if error != nil {
		return 0
	}
	
	defer db.Close()
	queryString := fmt.Sprintf("SELECT CodeID FROM code WHERE CodeName='%s'",qr)
	rows, error := db.Query(queryString)
	if error := rows.Err(); error != nil {
		return 0
	}
	var id int
	if rows.Next() {
		if error := rows.Scan(&id); error != nil {
			return 0	
		}
	}else{
		return 0
	}
	rows.Close()
	
	return id
}