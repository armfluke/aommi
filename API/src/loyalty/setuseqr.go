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
	if qr.Split(qr,"|")[1]=="saving_account"{
		rows, error := db.Query("UPDATE account set SavingAccount=1")
		if  error != nil {
			return false
		}
	}else if qr.Split(qr,"|")[1]=="fixed_account"{
		rows, error := db.Query("UPDATE account set FixedAccount=1")
		if  error != nil {
			return false
		}
	}

	idqr := GetIDQR(qr);
	queryString := fmt.Sprintf("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('%s',%d,%d)",account,idqr,point)
	rows, error := db.Query(queryString)
	if  error != nil {
		return false
	}
	rows.Close()
	
	return true
}