package main

import (
	"log"
	"loyalty"
	"net/http"
	"loyaltyWeb"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/account", loyalty.GetAccount)
	http.HandleFunc("/promotion", loyalty.GetPromotion)
	http.HandleFunc("/promotion/use", loyalty.UsePromotion)
	http.HandleFunc("/point/update", loyalty.UpdatePoint)
	http.HandleFunc("/web/account", loyaltyWeb.WebViewAccount)
	http.HandleFunc("/web/promotionused", loyaltyWeb.WebViewPromotionUsed)
	http.HandleFunc("/qr", loyalty.ScanQrCode)

	log.Println("Server running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
