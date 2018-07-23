package loyalty

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		message := Status{StatusCode: 500, Status: "Can't connect to database"}

		encoder := json.NewEncoder(w)
		encoder.Encode(message)
	}

	defer db.Close()

	queryString := fmt.Sprintf("UPDATE account SET PointBalance = %d WHERE AccountID = '%s'",
		decreasepoint.PointBalance,
		decreasepoint.AccountID)
	results, err := db.Query(queryString)
	results.Close()

	if err != nil {
		message := Status{StatusCode: 500, Status: "Can't insert to database"}

		encoder := json.NewEncoder(w)
		encoder.Encode(message)
	}

	message := Status{StatusCode: 200, Status: "ok"}

	encoder := json.NewEncoder(w)
	encoder.Encode(message)
}
