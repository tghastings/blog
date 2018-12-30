// Hastings Blog
//
// A custom blog API built on golang.
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Thomas Hastings<thomas@thomashastings.com> https://www.thomashastings.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tghastings/blog/admin/db"
	user "github.com/tghastings/blog/admin/users"
	post "github.com/tghastings/blog/posts"
)

var mySigningKey = []byte("pleasedonthackmebro")

var err error

func main() {
	if err := db.Open(); err != nil {
		// handle error
	}
	defer db.Close()
	// Migrate the schema
	db.DB.AutoMigrate(user.User{})
	db.DB.AutoMigrate(post.Post{})
	// Check to see if a user exisits if not create root
	user.FirstUser()

	/*
		ROUTES
	*/

	// swagger:route GET /posts posts
	//
	// Lists posts sorted by more recent post first
	//
	// This will show all available posts by default.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.HandleFunc("/", post.Index)

	// swagger:route GET /post/{post} posts
	//
	// Lists posts sorted by more recent post first
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.HandleFunc("/post", post.Show)

	// swagger:route POST /post/create posts
	//
	// Create a new post
	//
	// This will allow the user to create a new post
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/post/create", isAuthorized(post.Create))

	// swagger:route POST /post/delete/{post} posts
	//
	// Delete an exisiting post.
	//
	// This will allow the user to delete an exisiting post
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/post/delete", isAuthorized(post.Delete))

	// swagger:route POST /post/update/{post} posts
	//
	// Update an exisiting post.
	//
	// This will allow the user to Update an exisiting post
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/post/update", isAuthorized(post.Update))

	// Users

	// swagger:route GET /admin/users users
	//
	// See all registered users.
	//
	// This will allow the user to see all registered users.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/users", isAuthorized(user.Index))

	// swagger:route GET /admin/user/{user} users
	//
	// Show a specific users
	//
	// This will allow the user to view a specific user record
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/user", isAuthorized(user.Show))

	// swagger:route POST /admin/user/create users
	//
	// Allow a registered user to create an account
	//
	// This will allow the user to create a new account
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/user/create", isAuthorized(user.Create))

	// swagger:route POST /admin/user/delete/{user} users
	//
	// Delete an exisiting user.
	//
	// This will allow the user to delete an exisiting user
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/user/delete", isAuthorized(user.Delete))

	// swagger:route POST /admin/user/update users
	//
	// Update an exisiting user
	//
	// This will allow the user to update an exisiting user
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.Handle("/admin/user/update", isAuthorized(user.Update))

	//Auth

	// swagger:route POST /auth users
	//
	// Authenticate a user with username and password.
	//
	// This will allow a user to authenticate with the system and get a JWT.
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	http.HandleFunc("/auth", user.Auth)

	//Swagger
	fs := http.FileServer(http.Dir("./docs/swagger"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// Start the application
	fmt.Println("The application has started and is listening on port 8090.")
	http.ListenAndServe(":8090", nil)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
