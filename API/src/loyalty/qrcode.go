package loyalty

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	BalancePoint int `json:"balancePoint"`
	QrPoint int `json:"qrPoint"`
}

func ScanQrCode(w http.ResponseWriter, r *http.Request) {	
	
	bodyPoint := Message{BalancePoint: 150,QrPoint: 150}
	

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(bodyPoint)
}