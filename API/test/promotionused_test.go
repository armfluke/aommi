package promotionused_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Promotion_Json_Input_PromotionID_123_AccountID_1100400758552_Rewardcode_xxxyyy_Sholdbe_Json_Statuscode_200_Status_ok(t *testing.T) {
	expected := `{"statusCode":"200","status":"ok"}`
	msg := `{"promotionID":"123","accountID":"1100400758552,"rewardcode":"xxxyyy"}`

	r, _ := http.NewRequest(http.MethodPost, "/promotion-is-used", strings.NewReader(msg))
	w := httptest.NewRecorder()

	Promotion_Is_Use_Post(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}

}
