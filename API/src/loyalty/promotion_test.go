package loyalty_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Promotion_Should_Return_All_Promotion(t *testing.T) {
	expected := `{"promotionID":1,"promotionName":"ตั๋วหนังฟรีMajor","reward":"ตั๋วหนังMajor","condition":"","point":100,"image":""}`
	r, _ := http.NewRequest(http.MethodGet, "/promotion", nil)
	w := httptest.NewRecorder()

	GetPromotion(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}
}

func Test_Promotion_Json_Input_PromotionID_1_AccountID_1100400758552_Rewardcode_MAJ123ABC_Should_Be_Json_StatusCode_200_Status_ok(t *testing.T) {
	expected := `{"statusCode":200,"status":"ok"}`
	msg := `{"accountID":"1100400758552,"promotionID":1,"rewardCode":"MAJ123ABC"}`

	r, _ := http.NewRequest(http.MethodPost, "/promotion/use", strings.NewReader(msg))
	w := httptest.NewRecorder()

	UsePromotion(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}

}
