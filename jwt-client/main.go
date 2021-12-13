package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte("Labardudu")
func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/token", Token)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func Token(w http.ResponseWriter, r *http.Request) {
	token, err := getJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprintf(w, token)

}
func getJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}