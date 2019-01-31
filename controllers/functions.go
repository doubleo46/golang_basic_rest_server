package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"

	"time"
)

var tokenAuth *jwtauth.JWTAuth

func GetPasswordHash(password string) string{
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	return string(hash)
}

func IsItAValidUser(password string) bool{
	// Get password hasg from DB
	hashFromDatabase := []byte("Get password hash from database")
    // Comparing the password with the hash
    if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(password)); err != nil {
        // TODO: Properly handle error
		log.Fatal(err)
		return false
    }
	return true
}
	func GetJWTToken(username string) string{
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
		// JWT token is based username and current time.
		_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"username": username,"time": time. Now()})
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
		return tokenString
	}


