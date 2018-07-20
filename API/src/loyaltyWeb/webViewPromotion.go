package loyaltyWeb

import (
	"net/http"
)
func WebViewPromotion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	message := `<!DOCTYPE html>
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
	w.Write([]byte(message))
}
