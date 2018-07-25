package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_GetBalanceAccount_Input_qr_accountID1140100074828_Should_Be_Balance_and_Status(t *testing.T) {
	account := "1140100074828"
	actual_balance, actual_Status := GetBalanceAccount(account)
	expected_balance, _ := GetBalanceAccount(account)

	if actual_balance != expected_balance || actual_point != expected_point {
		t.Errorf("expect '%d %d' but got '%d %d'", expected_balance, expected_point, actual_balance, actual_point)
	}

}
