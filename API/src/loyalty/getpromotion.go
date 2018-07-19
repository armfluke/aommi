package loyalty

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

var promotion []Promotion

func GetPromotionFromDB(username string,
	password string,
	host string,
	database string) string {
	log.Println("Connected to database...")
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			username,
			password,
			host,
			database))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM promotion")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		var data Promotion
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			//checkColumns
			switch columns[i] {
			case "PromtotionID":
				data.PromotionID, _ = strconv.Atoi(value)
			case "PromotionName":
				data.PromotionName = value
			case "Reward":
				data.Reward = value
			case "Redeem_condition":
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
			// fmt.Println(columns[i], ": ", value)
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	arrPromotion, _ := json.Marshal(promotion)
	return string(arrPromotion)
}
