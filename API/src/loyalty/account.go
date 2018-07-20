package loyalty

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	AccountID    string `json:"accountID"`
	AccountName  string `json:"accountName"`
	PointBalance int    `json:"pointBalance"`
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := GetAccountFromDatabase()

	w.Write([]byte(message))
}

func GetAccountFromDatabase() string {
	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		return ""
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM account")
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

	var account []Account

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return ""
		}

		var value string
		var data Account
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i] {
			case "AccountID":
				data.AccountID = value
			case "AccountName":
				data.AccountName = value
			case "PointBalance":
				data.PointBalance, _ = strconv.Atoi(value)
			}

			if i == len(columns)-1 {
				account = append(account,
					Account{
						AccountID:    data.AccountID,
						AccountName:  data.AccountName,
						PointBalance: data.PointBalance,
					})
			}

			fmt.Println(columns[i] + " : " + value)
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrAccount, _ := json.Marshal(account)
	return string(arrAccount)
}
