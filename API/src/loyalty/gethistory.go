package loyalty

import (
	"encoding/json"
	"net/http"
)

type PostHistoryData struct {
	AccountID string `json:accountID`
}

type HistoryFunc struct {
	GetHistoryFromDatabase func(string) string
}

func (hist HistoryFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data PostHistoryData
	body := json.NewDecoder(r.Body)
	body.Decode(&data)

	message := hist.GetHistoryFromDatabase(data.AccountID)

	w.Write([]byte(message))
}
