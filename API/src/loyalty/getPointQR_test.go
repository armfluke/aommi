package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_GetPointQR_Input_Key_200savingaccount_Should_Be_Output_200(t *testing.T) {
	qrCode := "200|saving_account"
	expectedPoint := 200
	expectedStatus := true
	actualPoint, actualStatus := GetPointQR(qrCode)

	if actualPoint != expectedPoint || actualStatus != expectedStatus {
		t.Errorf("expect '%d %t' but got '%d %t'", expectedPoint, expectedStatus, actualPoint, actualStatus)
	}
}
