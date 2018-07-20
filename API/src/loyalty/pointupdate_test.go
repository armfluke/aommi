package loyalty_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Promotion_Json_Input_AccountID_1100400758552_PointBalance_900_Should_Be_Json_StatusCode_200_Status_ok(t *testing.T) {
	expected := `{"statusCode":200,"status":"ok"}`
	message := `{"accountID":"1100400758552","pointBalance":900}`

	r, _ := http.NewRequest(http.MethodPost, "/point/update", strings.NewReader(message))
	w := httptest.NewRecorder()

	UpdatePoint(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}

}
