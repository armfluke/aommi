package loyalty

import (
	"fmt"
)

func GetIDQR(qr string) int{
	
	db := ConnectDatabase();
	if db==nil{
		return 0
	}
	
	defer db.Close()
	queryString := fmt.Sprintf("SELECT CodeID FROM code WHERE CodeName='%s'",qr)
	rows, _ := db.Query(queryString)
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