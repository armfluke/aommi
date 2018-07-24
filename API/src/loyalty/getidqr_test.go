package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_GetIDQR_Input_Key_200savingaccount_Should_Be_Output_200(t *testing.T) {
	qrCode := "200|saving_account"
	expectedResult := 1
	actualResult := GetIDQR(qrCode)

	if actualResult != expectedResult {
		t.Errorf("expect '%d' but got '%d'", expectedResult, actualResult)
	}
}
