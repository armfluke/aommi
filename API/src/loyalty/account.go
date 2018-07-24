package loyalty

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	AccountID    string `json:"accountID"`
	AccountName  string `json:"accountName"`
	PointBalance int    `json:"pointBalance"`
}

type Customer struct {
	GetAccountFromDatabase func(string) string
}
func (customer Customer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["accountID"]
	if keys != nil {
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		key := keys[0]
		message := customer.GetAccountFromDatabase(key)
		w.Write([]byte(message))
	} else {
		message := GetAllAccountFromDatabase()
		w.Write([]byte(message))
	}
}



func GetAllAccountFromDatabase() string {
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
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrAccount, _ := json.Marshal(account)
	return string(arrAccount)
}

func GetAccountFromDatabase(key string) string {
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
				if data.AccountID == key {
					account = append(account,
						Account{
							AccountID:    data.AccountID,
							AccountName:  data.AccountName,
							PointBalance: data.PointBalance,
						})
				}
			}
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrAccount, _ := json.Marshal(account)
	return string(arrAccount)
}

func GetAccountForTest(w http.ResponseWriter, r *http.Request) {
	message := `[{"accountID":"1100400758552","accountName":"ปวรืศร มยานนท์","pointBalance":0},{"accountID":"1140100074828","accountName":"วรพรต เดชลรัตน์","pointBalance":3000}]`
	w.Header().Set("Content-Type", "application/json")	
	w.Write([]byte(message))
}
