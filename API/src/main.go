package main

import (
	"log"
	"loyalty"
	"loyaltyWeb"
	"net/http"
)

func main() {
	http.HandleFunc("/account", loyalty.GetAccount)
	http.HandleFunc("/promotion", loyalty.GetPromotion)
	http.HandleFunc("/promotion/use", loyalty.UsePromotion)
	http.HandleFunc("/web/promotion", loyaltyWeb.WebViewPromotion)
	log.Println("Server running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
