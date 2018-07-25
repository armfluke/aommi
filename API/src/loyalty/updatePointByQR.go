package loyalty

func UpdatePointByQR(qr string, account string) (int, int, bool) {

	pointqr, status := GetPointQR(qr)
	resultbalance := 0
	if !status {
		return 0, 0, false
	}
	accountBalance, status := GetBalanceAccount(account)
	if !status {
		return 0, 0, false
	}
	status = IsUseQR(account, qr)
	if !status {
		return 0, 0, false
	}
	status = SetUseQR(account, qr, pointqr)
	if !status {
		return 0, 0, false
	}

	resultbalance = accountBalance + pointqr
	if UpdatePointToDatabase(resultbalance, account) {
		return resultbalance, pointqr, true
	} else {
		return 0, 0, false
	}
}
