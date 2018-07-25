package loyalty_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetHistory_Json_Input_AccountID_1100400758552_Should_Be_HistoryData(t *testing.T) {
	expected := `{"statusCode":200,"status":"ok"}`
	msg := `{"accountID":"1100400758552"}`

	r, _ := http.NewRequest(http.MethodPost, "/account/history", strings.NewReader(msg))
	w := httptest.NewRecorder()

	GetHistory(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}

}
