package loyaltyWeb

import (
	"net/http"
	"database/sql"
	"bytes"
	"fmt"
)

type Account struct {
	AccountID   string    
	AccountName string 
	PointBalance string   
}

func WebViewAccount(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	database, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")
	if error != nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()

	rows, error := database.Query("SELECT * FROM account")
	columns, error := rows.Columns()
	if error != nil {
		return 
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var message_account bytes.Buffer
	for rows.Next() {
		error = rows.Scan(scanArgs...)
		if error != nil {
			return 
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
				data.PointBalance = value
			}

			if i == len(columns)-1 {
				
				message_account.WriteString(`<tr>
					<td>`+data.AccountID+`</td><td>`+data.AccountName+`</td><td>`+data.PointBalance+`</td>
				</tr>`)
			}
		}
	}

	if error = rows.Err(); error != nil {
		return 
	}
	var message bytes.Buffer
	message.WriteString(`
	 <html>
	 <body>
	 <table>
	   <tr>
	 	<th>AccountID</th>
	 	<th>AccountName</th>
	 	<th>PointBanlance</th>
	   </tr>`)
	   message.WriteString(message_account.String())
	   message.WriteString(`</table>
	 </body>
	 </html>
	 `)
	w.Write([]byte(message.String()))
}
