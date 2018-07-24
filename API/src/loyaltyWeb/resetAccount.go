package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
	"log"
)

func ResetAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                     // Parses the request body
    
	tmpl := template.Must(template.ParseFiles("loyaltyWeb/webViewAccount.html"))
	database := ConnectDatabase()
	if database == nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()
	typeAccount := r.Form.Get("typeaccount")
	accountReset := r.Form.Get("account")
	queryString := ""
	if typeAccount=="Saving"{
		queryString = "UPDATE account set SavingAccount=0 WHERE AccountID=?"
	}else if typeAccount=="Fixed"{
		queryString = "UPDATE account set FixedAccount=0 WHERE AccountID=?"
	} 
	_,error := database.Exec(queryString,accountReset)
	if error != nil {
		log.Println(error)
		fmt.Fprintf(w,"Error Reset")
		return
	}
	
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
	