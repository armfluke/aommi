package loyalty

import (
	"encoding/json"
	"net/http"
)

type PostBodyQR struct {
	QRCode    string `json:"qrCode"`
	AccountID string `json:"accountID"`
}

type MessageQR struct {
	BalancePoint int `json:"balancePoint"`
	QrPoint      int `json:"qrPoint"`
}

func ScanQrCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var postBodyQR PostBodyQR
	var responseQR MessageQR

	body := json.NewDecoder(r.Body)
	body.Decode(&postBodyQR)

	//to do
	checkStatus := CheckQRCode(postBodyQR.QRCode)
	if checkStatus {
		resultBalance, resultQRPoint, status := UpdatePointByQR(postBodyQR.QRCode,
			postBodyQR.AccountID)
		if status {
			responseQR = MessageQR{BalancePoint: resultBalance, QrPoint: resultQRPoint}
		}
	} else {
		responseQR = MessageQR{BalancePoint: 0, QrPoint: 0}
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(responseQR)
}
