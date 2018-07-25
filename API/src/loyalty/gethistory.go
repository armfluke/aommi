package loyalty

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostHistoryData struct {
	AccountID string `json:accountID`
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data PostHistoryData
	body := json.NewDecoder(r.Body)
	body.Decode(&data)

	fmt.Println(data.AccountID)

	message := GetHistoryFromDatabase(data.AccountID)

	w.Write([]byte(message))
}
