package main

import (
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
}
