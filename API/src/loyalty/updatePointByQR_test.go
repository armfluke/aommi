package loyalty_test

import (
	"testing"
	. "loyalty"
)

func Test_UpdatePointByQR_Input_qr_200savingaccount_and_account1969900224728_Should_Be_Banlance_and_point(t *testing.T) {
	qrCode := "200|saving_account"
	account := "1969900224728"
	actual_balance,actual_point,_ := UpdatePointByQR(qrCode,account)
	balance,_ := GetBalanceAccount(account) 
	point,_ := GetPointQR(qrCode)
	expected_balance,expected_point := balance,point  
	
	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance,expected_point, actual_balance,actual_point)
	}

}
