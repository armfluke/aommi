package loyalty

import (
	
)

func UpdatePointByQR(qr string, account string) (int,int,bool){

	pointqr,error := GetPointQR(qr)
	resultbalance := 0
	if !error {
		return 0,0,false
	}
	accountbalance,error := GetBalanceAccount(account)
	if !error {
		return 0,0,false
	}
	if accountbalance>=pointqr {
		resultbalance = accountbalance-pointqr
	}
	if updatePointToDatabase(resultbalance,account){
		return accountbalance,pointqr,true
	}else {
		return 0,0,false
	}
}