package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("pleasedonthackmebro")

// IsAuthorized checks to see if a user has a current JWT
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		fmt.Print("Here you are!")
		if r.Header["Authorization"] != nil {
			fmt.Println("Here is the token auth: ")
			fmt.Println(r.Header["Authorization"][0])
			fmt.Println("Here you are!")
			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Println("error here!")
				fmt.Println(err)
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				fmt.Println("Authorized")
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// GenerateJWT generates the JSON web token
func GenerateJWT(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = username

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "Error, unable to make JWT in user.go"
	}
	//json resp
	return tokenString
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}