package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
)

type CustomersRedeemCode struct {
	AccountName   string    
	Reward string 
	Point string
	DateUsed string  
}
type DataToPageRedeem struct {
	CustomerList []CustomersRedeemCode 
}

func GetAccountRedeem(w http.ResponseWriter, r *http.Request) {
	
	tmpl := template.Must(template.ParseFiles("loyaltyWeb/getaccountredeem.html"))
	database := ConnectDatabase()
	if database == nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()
	queryString := "SELECT account.AccountName,promotion.Reward,promotion.Point,Date(promotionUsed.DateUsed) FROM promotionUsed INNER JOIN promotion ON promotion.PromotionID=promotionUsed.PromotionID INNER JOIN account ON account.AccountID=promotionUsed.AccountID ORDER BY promotionUsed.DateUsed DESC"
	rows, error := database.Query(queryString)
	if error != nil {
		fmt.Fprintf(w,"Error Query")
		return
	}
	data := DataToPageRedeem{}
	for rows.Next() {
		var cdata CustomersRedeemCode
		error = rows.Scan(&cdata.AccountName, &cdata.Reward, &cdata.Point, &cdata.DateUsed)
		if error != nil {
			fmt.Fprintf(w,"Error Query save")
			return
		}
		data.CustomerList = append(data.CustomerList,cdata)
	}
	tmpl.Execute(w, data)
}
	