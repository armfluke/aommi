package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_CheckQRCode_Input_qrCode_200savingaccount_Should_Be_Status_True(t *testing.T) {
	qrCode := "200|saving_account"
	expectedResult := true
	actualResult := CheckQRCode(qrCode)

	if actualResult != expectedResult {
		t.Errorf("expect '%t' but got '%t'", expectedResult, actualResult)
	}
}

func Test_GetQRCodeFromDatabase_Input_Key_200savingaccount_Should_Be_Output_200saving_account(t *testing.T) {
	expectedResult := "200|saving_account"
	actualResult := GetQRCodeFromDatabase(expectedResult)

	if actualResult != expectedResult {
		t.Errorf("expect '%s' but got '%s'", expectedResult, actualResult)
	}
}
