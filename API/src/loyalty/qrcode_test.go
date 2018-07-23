package loyalty_test

import (
	"io/ioutil"
	. "loyalty"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Scan_QR_Code(t *testing.T){
	expectedResult := `{"balancePoint":150,"qrPoint":150}`
	message := `{"qrCode":"xxxyyy","accountID":"1100400758552"}`

	r, _ := http.NewRequest(http.MethodPost, "/qr", strings.NewReader(message))
	w := httptest.NewRecorder()

	ScanQrCode(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actualResult := strings.TrimSpace(string(body))

	if actualResult != expectedResult {
		t.Errorf("expect '%s' but got '%s'", expectedResult, actualResult)
	}
}