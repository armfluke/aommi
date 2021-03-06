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
	expected := `[{"promotionID":1,"promotionName":"ตั๋วหนัังฟรีMajor","reward":"ตั๋วหนังMajor","condition":"","point":100,"image":""},{"promotionID":2,"promotionName":"Starbuck voucher","reward":"Starbuck voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":3,"promotionName":"Digital Salak","reward":"Digital Salak 1 unit","condition":"","point":100,"image":""},{"promotionID":4,"promotionName":"Major movie ticket","reward":"Major movie 1 unit","condition":"","point":300,"image":""},{"promotionID":5,"promotionName":"SF movie ticket","reward":"SF movie 1 unit","condition":"","point":300,"image":""},{"promotionID":6,"promotionName":"True Coffee voucher","reward":"True Coffee voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":7,"promotionName":"Central gift voucher","reward":"Central gift voucher 100 bath","condition":"","point":200,"image":""},{"promotionID":8,"promotionName":"Bag","reward":"Bag","condition":"","point":400,"image":""}]`
	r, _ := http.NewRequest(http.MethodGet, "/promotion", nil)
	w := httptest.NewRecorder()

	stubGetPromotionFromDatabase := func() string {
		return expected
	}

	product := Product{stubGetPromotionFromDatabase}

	product.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect\n'%s' but got\n'%s'", expected, actual)
	}
}

func Test_Promotion_Json_Input_PromotionID_1_AccountID_1100400758552_Rewardcode_MAJ123ABC_Should_Be_Json_StatusCode_200_Status_ok(t *testing.T) {
	expected := `{"statusCode":200,"status":"ok"}`
	msg := `{"accountID":"1100400758552","promotionID":1,"rewardCode":"MAJ123ABC"}`

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
