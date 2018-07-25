package loyalty

import (
	_ "github.com/go-sql-driver/mysql"
)

type QRCode struct {
	CodeName string `json:"CodeName"`
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

	rows, err := db.Query("SELECT CodeName FROM code WHERE CodeName=?", key)
	if err != nil {
		return ""
	}

	var qrcode []QRCode
	var qr QRCode
	for rows.Next() {
		err = rows.Scan(&qr.CodeName)
		if err != nil {
			return ""
		}
		qrcode = append(qrcode, qr)
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
