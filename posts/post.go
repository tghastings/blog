package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tghastings/blog/admin/db"
)

type Post struct {
	gorm.Model
	Title   string
	Date    string
	Tags    string
	Author  string
	Content string `gorm:"type:varchar(256)"`
}

type Message struct {
	Type    string
	Message string
}

func Index(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	db.DB.Order("created_at desc").Find(&posts)
	js, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /admin/post/create posts
	//
	// Create new post
	//
	//
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// parameters:
	// - name: tags
	//   in: query
	//   description: tags to filter by
	//   required: false
	//   type: array
	//   items:
	//     type: string
	//   collectionFormat: csv
	// - name: limit
	//   in: query
	//   description: maximum number of results to return
	//   required: false
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: pet response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/pet"
	//   default:
	//     description: unexpected error
	//     schema:
	//       "$ref": "#/definitions/errorModel"
	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		// handle error
		log.Fatal("Error")
	}
	if err != nil {
		log.Fatal(err)
	}
	// Create
	db.DB.Create(&newPost)

	//json resp
	msg := Message{"success", "Post added."}
	js, err := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Show(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(post.ID)
	db.DB.Find(&post, post.ID)
	js, err := json.Marshal(&post)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(post)
	db.DB.Model(&post).Where("ID = ?", post.ID).Updates(post)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(post.ID)
	// unmarshal content to ApplicantJSON
	db.DB.Find(&post, post.ID)
	db.DB.Delete(&post)
}
