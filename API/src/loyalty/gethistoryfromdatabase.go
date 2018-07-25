package loyalty

import (
	"encoding/json"
	"fmt"
)

type History struct {
	Image  string `json:"image"`
	Reward string `json:"reward"`
	Point  int    `json:"point"`
	Code   string `json:"code"`
	Date   string `json:"date"`
}

func GetHistoryFromDatabase(account string) string {

	db := ConnectDatabase()
	if db == nil {
		return ""
	}
	defer db.Close()

	queryString := fmt.Sprintf("SELECT promotion.Image,promotion.Reward,promotion.Point,promotionUsed.RewardCode,promotionUsed.DateUsed FROM promotionUsed INNER JOIN promotion ON promotion.PromotionID=promotionUsed.PromotionID WHERE promotionUsed.AccountID='%s' ORDER BY promotionUsed.DateUsed DESC", account)
	rows, _ := db.Query(queryString)
	if err := rows.Err(); err != nil {
		return ""
	}
	var history []History
	for rows.Next() {
		var transaction History
		if error := rows.Scan(&transaction.Image, &transaction.Reward, &transaction.Point, &transaction.Code, &transaction.Date); error != nil {
			return ""
		}
		history = append(history, transaction)
	}
	rows.Close()

	arrHistory, _ := json.Marshal(history)
	return string(arrHistory)
}
