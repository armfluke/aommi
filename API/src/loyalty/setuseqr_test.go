package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_SetUseQR_Input_account_1140100074828_and_qr_200savingaccount_and_point_200_Should_Be_Stauts(t *testing.T) {
	account := "1140100074828"
	qr := "200|saving_account"
	point := 200

	expected_status := true
	actual_status := SetUseQR(account, qr, point)

	if actual_status != expected_status {
		t.Errorf("expect '%t' but got '%t'", expected_status, actual_status)
	}

}
