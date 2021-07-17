package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dustin-ward/recordrank/db"
	"github.com/dustin-ward/recordrank/routes"
)

func main() {
	fmt.Println("Starting API Service...")
	fmt.Println("Connecting to DB...")
	db.Connect()
	defer db.Disconnect()

	fmt.Println("Configuring Routes...")
	router := routes.Setup()

	fmt.Println("API Service Running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
