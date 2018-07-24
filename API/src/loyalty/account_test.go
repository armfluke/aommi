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
	expected := `[{"accountID":"1100400758552","accountName":"ปวรืศร มยานนท์","pointBalance":0},{"accountID":"1140100074828","accountName":"วรพรต เดชลรัตน์","pointBalance":3000}]`
	r, _ := http.NewRequest(http.MethodGet, "/account?accountID=1", nil)
	w := httptest.NewRecorder()

	stubGetAccountFromDatabase := func(key string) string{
		return expected
	}

	customer := Customer{stubGetAccountFromDatabase}

	customer.ServeHTTP(w,r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect\n'%s' but got\n'%s'", expected, actual)
	}
}

func Test_ActionToTest_Input_JSON_Should_Be_AllAccountJSON(t *testing.T) {
	expectedResult := `[{"accountID":"1100400758552","accountName":"ปวรืศร มยานนท์","pointBalance":0},{"accountID":"1140100074828","accountName":"วรพรต เดชลรัตน์","pointBalance":3000}]`

	actualResult := GetAllAccountFromDatabase()

	if expectedResult != actualResult {
		t.Errorf("expected\n%s but got\n%s", expectedResult, actualResult)
	}
}

func Test_ActionToTest_Input_JSON_Should_Be_AccountJSON(t *testing.T) {
	expectedResult := `[{"accountID":"1140100074828","accountName":"วรพรต เดชลรัตน์","pointBalance":3000}]`

	actualResult := GetAccountFromDatabase("1140100074828")

	if expectedResult != actualResult {
		t.Errorf("expected\n%s but got\n%s", expectedResult, actualResult)
	}
}
