package loyalty_test

import (
	. "loyalty"
	"testing"
)

func Test_Status_Update_Point_To_Database_Input_Point_Balance_200_And_Account_ID_1100400758552_ShuldBe_Return_Status_Ture(t *testing.T) {
	expactedResult := true

	pointBalance := 200
	accountID := "1100400758552"

	actualResult := UpdatePointToDatabase(pointBalance, accountID)

	if actualResult != expactedResult {
		t.Errorf("expect '%t' but got '%t'", expactedResult, actualResult)
	}
}
