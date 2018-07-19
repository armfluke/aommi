package loyalty_test

import (
	"loyalty"
	"testing"
)

func Test_ActionToTest_Input_JSON_Should_Be_PromotionJSON(t *testing.T) {
	expectedResult := `[{"promotionID":0,"promotionName":"ตั๋วหนัังฟรีMajor","reward":"ตั๋วหนังMajor","condition":"","point":100,"image":""}]`

	username := "root"
	password := "Admin123!"
	host := "178.128.48.140"
	database := "aommi"

	actualResult := loyalty.GetPromotionFromDB(username,
		password,
		host,
		database)

	if expectedResult != actualResult {
		t.Errorf("expected %s but got %s", expectedResult, actualResult)
	}
}
