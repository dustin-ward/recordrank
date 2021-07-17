package controllers

import (
	"fmt"
	"net/http"
)

func HomePagePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /")
	fmt.Fprintf(w, "RecordRank API: POST")
}

func HomePageGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /")
	fmt.Fprintf(w, "RecordRank API: GET")
}
