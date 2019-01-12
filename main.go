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

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	auth "github.com/tghastings/blog/api/auth"
	post "github.com/tghastings/blog/api/posts"
	user "github.com/tghastings/blog/api/users"
	"github.com/tghastings/blog/config/db"
)

var err error

func main() {
	if err := db.Open(); err != nil {
		// handle error
	}
	defer db.Close()
	// Migrate the schema
	db.DB.AutoMigrate(user.User{})
	db.DB.AutoMigrate(post.Post{})
	// Check to see if a user exists if not create root
	user.FirstUser()

	/*
		SWAGGER OPERATIONS
	*/

	// swagger:operation GET / posts
	//
	// Returns all posts from the system
	//
	// Returns all posts
	//
	// ---
	// produces:
	// - application/json
	//
	http.HandleFunc("/", post.Index)

	// swagger:operation GET /post/{id} posts
	//
	// Returns a single post from the system
	//
	// Returns a single post
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: body
	//   description: The ID of the post to retrieve
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	http.HandleFunc("/post/", post.Show)

	// swagger:operation POST /admin/post/ posts
	//
	// Creates a new post in the system
	//
	// Creates a new post
	//
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Title
	//   in: body
	//   description: The title of the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Date
	//   in: body
	//   description: The date of the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Author
	//   in: body
	//   description: The author for the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Content
	//   in: body
	//   description: The content of the post
	//   required: true
	//   type: object
	//   items:
	//     type: text
	//   collectionFormat: json
	// - name: Tags
	//   in: body
	//   description: The tags of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json

	// swagger:operation PUT /admin/post/{id} posts
	//
	// Update an existing post
	//
	// Updates an existing post
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: query
	//   description: The ID of the post to update
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Title
	//   in: body
	//   description: The title of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Date
	//   in: body
	//   description: The date of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Author
	//   in: body
	//   description: The author for the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Content
	//   in: body
	//   description: The content of the post
	//   required: false
	//   type: object
	//   items:
	//     type: text
	//   collectionFormat: json
	// - name: Tags
	//   in: body
	//   description: The tags of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json

	// swagger:operation DELETE /admin/post/{id} posts
	//
	// Deletes a single post from the system
	//
	// Deletes a single post
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: body
	//   description: The ID of the post to retrieve
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json

	http.Handle("/admin/post/", auth.IsAuthorized(post.Route))
	http.Handle("/admin/post", auth.IsAuthorized(post.Route))

	// Users

	// swagger:operation GET /admin/users users
	//
	// Returns all users from the system
	//
	// Returns all users
	//
	// ---
	// produces:
	// - application/json
	http.Handle("/admin/users", auth.IsAuthorized(user.Index))

	// swagger:operation POST /admin/user/ user
	//
	// Creates a new user in the system
	//
	// Creates a new user
	//
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Username
	//   in: body
	//   description: The title of the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Password
	//   in: body
	//   description: The title of the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Email Address
	//   in: body
	//   description: The title of the post
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json

	// swagger:operation PUT /admin/user/{id} user
	//
	// Updates an existing user in the system
	//
	// Updates an existing user
	//
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: query
	//   description: The ID of the post to retrieve
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Username
	//   in: body
	//   description: The title of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Password
	//   in: body
	//   description: The title of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Email Address
	//   in: body
	//   description: The title of the post
	//   required: false
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json

	// swagger:operation DELETE /admin/user/{id} posts
	//
	// Deletes a single user from the system
	//
	// Deletes a single user
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   in: body
	//   description: The ID of the post to retrieve
	//   required: true
	//   type: object
	//   items:
	//     type: string
	http.Handle("/admin/user/", auth.IsAuthorized(user.Route))
	http.Handle("/admin/user", auth.IsAuthorized(user.Route))

	//Auth

	// swagger:operation POST /auth auth
	//
	// Authenticates a user against the system
	//
	// Authenticates a user and returns a Token in the form of a cookie
	// ---
	// produces:
	// - headers/cookie
	// parameters:
	// - name: Username
	//   in: body
	//   description: The username of the user to authenticate
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	// - name: Password
	//   in: body
	//   description: The plain text password of the user to authenticate
	//   required: true
	//   type: object
	//   items:
	//     type: string
	//   collectionFormat: json
	http.HandleFunc("/auth", user.UserAuth)

	//Swagger
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// Start the application
	fmt.Println("The application has started and is listening on port 8090.")
	http.ListenAndServe(":8090", nil)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
