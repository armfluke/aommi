package loyalty

func UpdatePointByQR(qr string, account string) (int, int, bool) {

	pointqr, status := GetPointQR(qr)
	resultbalance := 0
	if !status {
		return 0, 0, false
	}
	accountbalance, status := GetBalanceAccount(account)
	if !status {
		return 0, 0, false
	}
	resultbalance = accountbalance + pointqr
	if updatePointToDatabase(resultbalance, account) {
		return resultbalance, pointqr, true
	} else {
		return 0, 0, false
	}
}
