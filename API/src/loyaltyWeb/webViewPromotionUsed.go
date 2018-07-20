package loyaltyWeb

import (
	"net/http"
	"database/sql"
	"bytes"
	"fmt"
)

type PromotionUsed struct {
	AccountName   string    
	PromotionName string 
	Point string
	DateUsed  string
}

func WebViewPromotionUsed(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	database, error := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")
	if error != nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()

	rows, error := database.Query("SELECT account.AccountName,promotion.PromotionName,promotion.Point,promotionUsed.DateUsed FROM promotionUsed promotionUsed inner join account account on account.AccountID=promotionUsed.AccountID inner join promotion promotion on promotionUsed.PromotionID=promotion.PromotionID")
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
		var data PromotionUsed
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i] {
			case "AccountName":
				data.AccountName = value
			case "PromotionName":
				data.PromotionName = value
			case "Point":
				data.Point = value
			case "DateUsed":
				data.DateUsed = value
			}

			if i == len(columns)-1 {
				
				message_account.WriteString(`<tr>
					<td>`+data.AccountName+`</td><td>`+data.PromotionName+`</td><td>`+data.Point+`</td><td>`+data.DateUsed+`</td>
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
	 <h3>Reprot History</h3>
	 <table>
	   <tr>
	 	<th>Account</th>
	 	<th>Promotion</th>
		 <th>Point</th>
		 <th>DataUsed</th>
	   </tr>`)
	   message.WriteString(message_account.String())
	   message.WriteString(`</table>
	 </body>
	 </html>
	 `)
	w.Write([]byte(message.String()))
}
