package loyalty

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type PointUpdate struct {
	AccountID    string `json:"accountID"`
	PointBalance int    `json:"pointBalance"`
}

func UpdatePoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var decreasepoint PointUpdate
	body := json.NewDecoder(r.Body)
	body.Decode(&decreasepoint)

	if (updatePointToDatabase(decreasepoint.PointBalance,decreasepoint.AccountID)){
		message := Status{StatusCode: 200, Status: "ok"}
		encoder := json.NewEncoder(w)
		encoder.Encode(message)
	}else{
		message := Status{StatusCode: 500, Status: "Can't connect to database"}

		encoder := json.NewEncoder(w)
		encoder.Encode(message)
	}
}
