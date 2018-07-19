package test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Promotion_Should_Return_All_Promotion(t *testing.T) {
	expected := `{"promotionID":0,"promotionName":"","reward":"","condition":"","point":0,"image":""}`
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
