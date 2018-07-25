package main

import (
	"log"
	"loyalty"	
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	http.Handle("/account", loyalty.Customer{loyalty.GetAccountFromDatabase})
	http.HandleFunc("/account/history", loyalty.GetHistory)
	http.Handle("/promotion", loyalty.Product{loyalty.GetPromotionFromDatabase})
	http.HandleFunc("/promotion/use", loyalty.UsePromotion)
	http.HandleFunc("/point/update", loyalty.UpdatePoint)
	
	http.HandleFunc("/qr", loyalty.ScanQrCode)
	

	log.Println("Server running on port 3001")

	log.Fatal(http.ListenAndServe(":3001", nil))	
}
