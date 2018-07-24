package loyalty

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	AccountID     string `json:"accountID"`
	AccountName   string `json:"accountName"`
	PointBalance  int    `json:"pointBalance"`
	SavingAccount int    `json:"savingAccount"`
	FixedAccount  int    `json:"fixedAccount"`
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
	db := ConnectDatabase()
	if db == nil {
		return ""
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM account")
	if err != nil {
		return ""
	}

	var account []Account
	var acc Account
	for rows.Next() {
		err = rows.Scan(&acc.AccountID, &acc.AccountName, &acc.PointBalance, &acc.SavingAccount, &acc.FixedAccount)
		if err != nil {
			return ""
		}
		account = append(account, acc)
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrAccount, _ := json.Marshal(account)
	return string(arrAccount)
}

func GetAccountFromDatabase(key string) string {
	db := ConnectDatabase()
	if db == nil {
		return ""
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM account")
	if err != nil {
		return ""
	}

	var account []Account
	var acc Account
	for rows.Next() {
		err = rows.Scan(&acc.AccountID, &acc.AccountName, &acc.PointBalance, &acc.SavingAccount, &acc.FixedAccount)
		if err != nil {
			return ""
		}

		if key == acc.AccountID {
			account = append(account, acc)
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	arrAccount, _ := json.Marshal(account)
	return string(arrAccount)
}
