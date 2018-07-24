package loyalty

import (
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
	db := ConnectDatabase()
	if db == nil {
		return ""
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM code")
	if err != nil {
		return ""
	}

	var qrcode []QRCode
	var qr QRCode
	for rows.Next() {
		err = rows.Scan(&qr.CodeID, &qr.CodeName, &qr.CodeType, &qr.CodePoint)
		if err != nil {
			return ""
		}

		if key == qr.CodeName {
			qrcode = append(qrcode, qr)
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
