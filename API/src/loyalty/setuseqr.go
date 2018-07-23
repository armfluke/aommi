package loyalty

import (
	"fmt"
)

func SetUseQR(account string, qr string,point int) bool{
	
	db := ConnectDatabase();
	if db==nil{
		return false
	}
	defer db.Close()
	
	idqr := GetIDQR(qr);
	queryString := fmt.Sprintf("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('%s',%d,%d)",account,idqr,point)
	rows, error := db.Query(queryString)
	if  error != nil {
		return false
	}
	rows.Close()
	
	return true
}