package loyalty

import (
	"database/sql"
	"fmt"
)

func GetPointQR(qr string) (int,bool){
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if error != nil {
		return 0,false
	}

	defer db.Close()
	queryString := fmt.Sprintf("SELECT CodePoint FROM code WHERE CodeName='%s'",qr)
	rows, error := db.Query(queryString)
	if error := rows.Err(); error != nil {
		return 0,false
	}
	var point int
	if rows.Next() {
		if error := rows.Scan(&point); error != nil {
			return 0,false	
		}
	}else{
		return 0,false
	}
	rows.Close()
	
	return point,true
}
