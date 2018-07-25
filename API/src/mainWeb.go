package main

import (
	"log"
	"loyaltyWeb"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("loyaltyWeb/"))))
	
	http.HandleFunc("/web/account", loyaltyWeb.WebViewAccount)
	http.HandleFunc("/web/account/earn", loyaltyWeb.GetAccountEarn)
	http.HandleFunc("/web/account/redeem", loyaltyWeb.GetAccountRedeem)
	http.HandleFunc("/web/qr", loyaltyWeb.GetQR)
	
	log.Println("Server running on port 8001")

	log.Fatal(http.ListenAndServe(":8001", nil))
}