package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
	"os" 
	"log"
)

type Customers struct {
	AccountID   string    
	AccountName string 
	PointBalance string
	SavingAccount bool
	FixedAccount bool    
}
type DataToPage struct {
	CustomerList []Customers 
}
func WebViewAccount(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println(dir)
	tmpl := template.Must(template.ParseFiles("loyaltyWeb/webviewaccount.html"))
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
		var cdata Customers
		var saving int
		var fixed int
		error = rows.Scan(&cdata.AccountID, &cdata.AccountName, &cdata.PointBalance, &saving, &fixed)
		if error != nil {
			fmt.Fprintf(w,"Error Query")
			return
		}
		if saving == 1 {
			cdata.SavingAccount = true
		}else{
			cdata.SavingAccount = false
		}
		if fixed == 1 {
			cdata.FixedAccount = true
		}else{
			cdata.FixedAccount = false
		}  
		data.CustomerList = append(data.CustomerList,cdata)
	}
	tmpl.Execute(w, data)
}
