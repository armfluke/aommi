package loyalty

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type QRCode struct {
	CodeID    int    `json:"CodeID"`
	CodeName  string `json:"CodeName"`
	CodeType  string `json:"CodeType"`
	CodePoint int    `json:"CodePoint"`
}

func CheckQRCode(qrCode string) bool {
	if qrCode != "" {
		message := GetQRCodeFromDatabase(qrCode)
		if message == qrCode {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func GetQRCodeFromDatabase(key string) string {
	db, err := sql.Open("mysql", "root:Admin123!@tcp(178.128.48.140:3306)/aommi")

	if err != nil {
		return ""
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM code")
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

	var qrcode []QRCode

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return ""
		}

		var value string
		var data QRCode
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i] {
			case "CodeID":
				data.CodeID, _ = strconv.Atoi(value)
			case "CodeName":
				data.CodeName = value
			case "CodeType":
				data.CodeType = value
			case "CodePoint":
				data.CodePoint, _ = strconv.Atoi(value)
				if data.CodeName == key {
					qrcode = append(qrcode,
						QRCode{
							CodeID:    data.CodeID,
							CodeName:  data.CodeName,
							CodeType:  data.CodeType,
							CodePoint: data.CodePoint,
						})
				}
			}
		}
	}

	if err = rows.Err(); err != nil {
		return ""
	}

	if qrcode != nil {
		return qrcode[0].CodeName
	} else {
		return ""
	}
}
