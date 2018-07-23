package loyalty_test

import (
	"testing"
	. "loyalty"
)

func Test_IsUseQR_Input_account1140100074828_and_qr_200savingaccount_Should_Be_True(t *testing.T) {
	qrCode := "200|saving_account"
	accountID := "1140100074828"
	actual := IsUseQR(accountID,qrCode)
	expected := true
	if actual != expected {
		t.Errorf("expect '%t' but got '%t'", expected, actual)
	}
	
}
func Test_IsUseQR_Input_account1140100074828_and_qr_200fixedgaccount_Should_Be_True(t *testing.T) {
	qrCode := "200|fixed_account"
	accountID := "1140100074828"
	actual := IsUseQR(accountID,qrCode)
	expected := true
	if actual != expected {
		t.Errorf("expect '%t' but got '%t'", expected, actual)
	}
	
}