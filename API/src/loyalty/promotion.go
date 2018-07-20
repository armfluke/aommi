package loyalty

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Promotion struct {
	PromotionID   int    `json:"promotionID"`
	PromotionName string `json:"promotionName"`
	Reward        string `json:"reward"`
	Condition     string `json:"condition"`
	Point         int    `json:"point"`
	Image         string `json:"image"`
}

type PromotionUsed struct {
	AccountID   string `json:"accountID"`
	PromotionID int    `json:"promotionID"`
	RewardCode  string `json:"rewardCode"`
}

type Status struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
}

func GetPromotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := GetPromotionFromDatabase()

	fmt.Println("Request Get")

	w.Write([]byte(message))
}

func UsePromotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var promotion PromotionUsed

	body := json.NewDecoder(r.Body)
	body.Decode(&promotion)

	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		message := Status{StatusCode: 500, Status: "Can't connect to database"}

		encoder := json.NewEncoder(w)
		encoder.Encode(message)
	}

	defer db.Close()

	results, err := db.Query("INSERT INTO promotionUsed (accountID, promotionID, rewardCode) VALUES ('" + promotion.AccountID + "'," + strconv.Itoa(promotion.PromotionID) + ",'" + promotion.RewardCode + "')")
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

func GetPromotionFromDatabase() string {
	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		return ""
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM promotion")
	if err != nil {
		return ""
	}

	columns, err := rows.Columns()
	if err != nil {
		return ""
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var promotion []Promotion

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return ""
		}

		var value string
		var data Promotion
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i] {
			case "PromotionID":
				data.PromotionID, _ = strconv.Atoi(value)
			case "PromotionName":
				data.PromotionName = value
			case "Reward":
				data.Reward = value
			case "RedeemCondition":
				data.Condition = value
			case "Point":
				data.Point, _ = strconv.Atoi(value)
			case "Image":
				data.Image = value
			}

			if i == len(columns)-1 {
				promotion = append(promotion,
					Promotion{
						PromotionID:   data.PromotionID,
						PromotionName: data.PromotionName,
						Reward:        data.Reward,
						Condition:     data.Condition,
						Point:         data.Point,
						Image:         data.Image,
					})
			}
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrPromotion, _ := json.Marshal(promotion)
	return string(arrPromotion)
}
