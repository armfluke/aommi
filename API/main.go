package main

import (
	"log"
	"loyalty"
	"net/http"
)

func main() {
	http.HandleFunc("/promotion", loyalty.GetPromotion)

	log.Println("Server running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
