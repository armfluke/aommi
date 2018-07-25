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
	expected := `[{"image":"https://scontent.fbkk8-1.fna.fbcdn.net/v/t1.0-1/p320x320/22894109_221818388356015_7936409751661051386_n.jpg?_nc_cat=0\u0026_nc_eui2=AeEWGQn_zRc0TUUtVqGiznn5mHkxOpmtTCXOFhFjtOzbJwWun-sJ1ugOBN6M-bHlPiIeEVVw-rQTn98r1d1hcLP_HIZkXzx_-bJNN8jEGvuuXw\u0026oh=c24b186755e6aaffc0dbbb01c27e73ec\u0026oe=5BD9CB0A","reward":"ตั๋วหนังMajor","point":100,"code":"MAJ123ABC","date":"2018-07-19 05:13:3}]`
	msg := `{"accountID":"1100400758552"}`

	r, _ := http.NewRequest(http.MethodPost, "/account/history", strings.NewReader(msg))
	w := httptest.NewRecorder()

	stubGetHistory := func(key string) string {
		return expected
	}
	hist := HistoryFunc{stubGetHistory}

	hist.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}

}
