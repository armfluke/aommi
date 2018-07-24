package loyaltyWeb

import (
	"net/http"
	"html/template"
	"fmt"
	. "loyalty"
)

type Code struct {
	Code   string    
	CodePoint string 
	Image string
}
type DataToPageCode struct {
	CodeList []Code 
}

func GetQR(w http.ResponseWriter, r *http.Request) {
	
	tmpl := template.Must(template.ParseFiles("loyaltyWeb/getqr.html"))
	database := ConnectDatabase()
	if database == nil {
		fmt.Fprintf(w,"Error Connect Database")
		return 
	}
	defer database.Close()
	queryString := "SELECT CodeType,CodePoint,Image FROM code"
	rows, error := database.Query(queryString)
	if error != nil {
		fmt.Fprintf(w,"Error Query")
		return
	}
	data := DataToPageCode{}
	for rows.Next() {
		var cdata Code
		error = rows.Scan(&cdata.Code, &cdata.CodePoint, &cdata.Image)
		if error != nil {
			fmt.Fprintf(w,"Error Query save")
			return
		}
		data.CodeList = append(data.CodeList,cdata)
	}
		
	tmpl.Execute(w, data)
}
	