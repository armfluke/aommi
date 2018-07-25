package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
)

type CustomersUseCode struct {
	AccountName   string    
	PointEarn int 
	DateEarn string
	CodeType string  
}
type DataToPageEarn struct {
	CustomerList []CustomersUseCode 
	SumAllPoints int
}

func GetAccountEarn(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("loyaltyWeb/getaccountearn.html"))
	database := ConnectDatabase()
	if database == nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()
	queryString := "SELECT account.AccountName,codeUsed.PointEarn,Date(codeUsed.DateEarn) AS DateEarn,code.CodeType FROM codeUsed INNER JOIN code ON code.CodeID=codeUsed.CodeID INNER JOIN account ON codeUsed.AccountID=account.AccountID ORDER BY codeUsed.DateEarn DESC"
	rows, error := database.Query(queryString)
	if error != nil {
		fmt.Fprintf(w,"Error Query")
		return
	}
	data := DataToPageEarn{}
	sumPoint := 0
	for rows.Next() {
		var cdata CustomersUseCode
		error = rows.Scan(&cdata.AccountName, &cdata.PointEarn, &cdata.DateEarn, &cdata.CodeType)
		if error != nil {
			fmt.Fprintf(w,"Error Query")
			return
		}
		sumPoint = sumPoint + cdata.PointEarn
		data.CustomerList = append(data.CustomerList,cdata)
	}
	data.SumAllPoints = sumPoint
	tmpl.Execute(w, data)
}
	