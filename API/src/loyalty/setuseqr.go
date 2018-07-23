package loyalty

import (
	"database/sql"
	"fmt"
)

func SetUseQR(account string, qr string,point int) bool{
	
	db, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if error != nil {
		return false
	}
	idqr := GetIDQR(qr);
	defer db.Close()
	queryString := fmt.Sprintf("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('%s',%d,%d)",account,idqr,point)
	rows, error := db.Query(queryString)
	if  error != nil {
		return false
	}
	rows.Close()
	
	return true
}