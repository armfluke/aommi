package loyaltyWeb_test

import (
	"io/ioutil"
	. "loyaltyWeb"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_PromotionByWeb_Should_Return_All_Promotion(t *testing.T) {
		expected := `<!DOCTYPE html>
	<html>
	<body>
	<table>
	  <tr>
		<th>AccountID</th>
		<th>AccountName</th>
		<th>PointBanlance</th>
	  </tr>
	  <tr>
		<td></td>
		<td></td>
		<td></td>
	  </tr>
	</table>
	
	</body>
	</html>`
	r, _ := http.NewRequest(http.MethodGet,"/web/account/",nil)
	w := httptest.NewRecorder()
	WebViewAccount(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))

	if actual != expected {
		t.Errorf("expect\n'%s' but got\n'%s'", expected, actual)
	}
}