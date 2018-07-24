package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
)

type Customer struct {
	AccountID   string    
	AccountName string 
	PointBalance string
	SavingAccount int
	FixedAccount int    
}
type DataToPage struct {
	customer []Customer 
}

func WebViewAccount(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("loyaltyWeb/webViewAccount.html"))
	database := ConnectDatabase()
	if database == nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()

	rows, error := database.Query("SELECT AccountID,AccountName,PointBalance,SavingAccount,FixedAccount FROM account")
	if error != nil {
		fmt.Fprintf(w,"Error Query")
		return
	}
	data := DataToPage{}
	for rows.Next() {
		var cdata Customer
		error = rows.Scan(&cdata.AccountID, &cdata.AccountName, &cdata.PointBalance, &cdata.SavingAccount, &cdata.FixedAccount)
		if error != nil {
			fmt.Fprintf(w,"Error Query")
			return
		}
		data.customer = append(data.customer,cdata)
	}
	fmt.Printf("%+v\n", data);
	tmpl.Execute(w, data)
}
