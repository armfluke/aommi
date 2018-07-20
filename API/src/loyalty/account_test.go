package loyalty_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Account_Should_Return_All_Account(t *testing.T) {
	expected := `[{"promotionID":0,"promotionName":"ตั๋วหนัังฟรีMajor","reward":"ตั๋วหนังMajor","condition":"","point":100,"image":""},{"promotionID":0,"promotionName":"Starbuck voucher","reward":"Starbuck voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Digital Salak","reward":"Digital Salak 1 unit","condition":"","point":100,"image":""},{"promotionID":0,"promotionName":"Major movie ticket","reward":"Major movie 1 unit","condition":"","point":300,"image":""},{"promotionID":0,"promotionName":"SF movie ticket","reward":"SF movie 1 unit","condition":"","point":300,"image":""},{"promotionID":0,"promotionName":"True Coffee voucher","reward":"True Coffee voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Central gift voucher","reward":"Central gift voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Bag","reward":"Bag","condition":"","point":400,"image":""}]`
	r, _ := http.NewRequest(http.MethodGet, "/account", nil)
	w := httptest.NewRecorder()

	GetAccount(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect\n'%s' but got\n'%s'", expected, actual)
	}
}

func Test_ActionToTest_Input_JSON_Should_Be_AccountJSON(t *testing.T) {
	expectedResult := `[{"promotionID":0,"promotionName":"ตั๋วหนัังฟรีMajor","reward":"ตั๋วหนังMajor","condition":"","point":100,"image":""},{"promotionID":0,"promotionName":"Starbuck voucher","reward":"Starbuck voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Digital Salak","reward":"Digital Salak 1 unit","condition":"","point":100,"image":""},{"promotionID":0,"promotionName":"Major movie ticket","reward":"Major movie 1 unit","condition":"","point":300,"image":""},{"promotionID":0,"promotionName":"SF movie ticket","reward":"SF movie 1 unit","condition":"","point":300,"image":""},{"promotionID":0,"promotionName":"True Coffee voucher","reward":"True Coffee voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Central gift voucher","reward":"Central gift voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":0,"promotionName":"Bag","reward":"Bag","condition":"","point":400,"image":""}]`

	actualResult := GetAccountFromDatabase()

	if expectedResult != actualResult {
		t.Errorf("expected\n%s but got\n%s", expectedResult, actualResult)
	}
}
