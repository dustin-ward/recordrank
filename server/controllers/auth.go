package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dustin-ward/recordrank/db"
	"github.com/dustin-ward/recordrank/util"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /auth/create")

	w.Header().Set("content-type", "application/json")

	var u db.User
	json.NewDecoder(r.Body).Decode(&u)

	// Email validation
	if !util.ValidEmail(u.Email) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(`{ "message": "user not found" }`))
		return
	}

	// Password Generation
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(password)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := db.GetCollection().InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /auth/login")

	w.Header().Set("content-type", "application/json")

	var u db.User
	json.NewDecoder(r.Body).Decode(&u)
	email := u.Email
	password := u.Password

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := db.GetCollection().FindOne(ctx, db.User{Email: email}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "user not found" }`))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{ "message": "incorrect password" }`))
		return
	}

	json.NewEncoder(w).Encode(u)
}
