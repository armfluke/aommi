package main

import (
<<<<<<< HEAD
	"fmt"
	"loyalty"
)

func main() {

	username := "root"
	password := "Admin123!"
	host := "178.128.48.140:3306"
	database := "aommi"

	results := loyalty.GetPromotionFromDB(
		username, password, host, database)
	fmt.Println(results)
=======
	"log"
	"loyalty"
	"net/http"
)

func main() {
	http.HandleFunc("/promotion", loyalty.GetPromotion)
	http.HandleFunc("/promotion/use", loyalty.UsePromotion)

	log.Println("Server running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
>>>>>>> 9e327dd8defe5c465b51a0e969af19e4e5b22eae
}
