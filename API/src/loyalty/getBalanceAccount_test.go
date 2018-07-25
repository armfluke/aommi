package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_GetBalanceAccount_Input_qr_accountID1140100074828_Should_Be_Balance_and_Status(t *testing.T) {
	account := "1140100074828"
	actual_balance, actual_Status := GetBalanceAccount(account)
	expected_balance := 650
	expected_status := true

	if actual_balance != expected_balance || actual_Status != expected_status {
		t.Errorf("expect '%d %t' but got '%d %t'", expected_balance, expected_status, actual_balance, actual_Status)
	}

}
