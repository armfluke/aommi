package loyalty

import (
	"fmt"
)

func GetPointQR(qr string) (int,bool){
	db := ConnectDatabase();
	if db==nil{
		return 0,false
	}

	defer db.Close()
	queryString := fmt.Sprintf("SELECT CodePoint FROM code WHERE CodeName='%s'",qr)
	rows, _ := db.Query(queryString)
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
