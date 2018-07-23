package loyalty_test

import (
	"testing"
	. "loyalty"
)

func Test_UpdatePointByQR_Input_qr_200savingaccount_and_account1140100074828_Should_Be_Banlance_and_point(t *testing.T) {
	qrCode := "200|saving_account"
	account := "1140100074828"
	actual_balance,actual_point,_ := UpdatePointByQR(qrCode,account)
	expected_balance,_ := GetBalanceAccount(account) 
	expected_point,_ := GetPointQR(qrCode)
	
	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance,expected_point, actual_balance,actual_point)
	}
	
}
func Test_UpdatePointByQR_Input_qr_200fixedgaccount_and_account1969900224728_Should_Be_0_and_0(t *testing.T) {
	qrCode := "200|fixed_account"
	account := "1969900224728"
	actual_balance,actual_point,_ := UpdatePointByQR(qrCode,account)
	expected_balance := 0 
	expected_point := 0
	
	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance,expected_point, actual_balance,actual_point)
	}
	
}
func Test_UpdatePointByQR_Input_qr_20deposit_and_account1140100074828_Should_Be_Balance_and_Point(t *testing.T) {
	qrCode := "20|deposit"
	account := "1140100074828"
	actual_balance,actual_point,_ := UpdatePointByQR(qrCode,account)
	expected_balance,_ := GetBalanceAccount(account) 
	expected_point,_ := GetPointQR(qrCode)
	
	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance,expected_point, actual_balance,actual_point)
	}
	
}
func Test_UpdatePointByQR_Input_qr_200deposit_and_account1140100074828_Should_Be_Balance_and_Point(t *testing.T) {
	qrCode := "200|deposit"
	account := "1140100074828"
	actual_balance,actual_point,_ := UpdatePointByQR(qrCode,account)
	expected_balance := 0 
	expected_point := 0
	
	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance,expected_point, actual_balance,actual_point)
	}
	
}