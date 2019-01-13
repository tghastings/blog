package post

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/tghastings/blog/config/db"
)

// Post describes the database schema
type Post struct {
	gorm.Model
	Title   string
	Date    string `gorm:"type:varchar(256)"`
	Tags    string
	Author  string
	Content string `gorm:"type:varchar(256)"`
}

// Message describes the JSON object message
type Message struct {
	Type    string
	Message string
}

// Index shows all posts
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

// Route is used to route methods for individual user actions
func Route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Show(w, r)
	case http.MethodPost:
		// Create a new record.
		Create(w, r)
	case http.MethodPut:
		// Update an existing record.
		Update(w, r)
	case http.MethodDelete:
		// Remove the record.
		Delete(w, r)
	default:
		http.Error(w, "Error with post routing", 400)
		return
	}
}

// Create new post
func Create(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Show displays a single post
func Show(w http.ResponseWriter, r *http.Request) {
	var post Post
	postID := strings.TrimPrefix(r.URL.Path, "/post/")
	if postID == "" {
		log.Println("ID Param 'key' is missing")
		return
	}
	db.DB.Find(&post, postID)
	js, err := json.Marshal(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Update updates one record
func Update(w http.ResponseWriter, r *http.Request) {
	var post Post
	postID := strings.TrimPrefix(r.URL.Path, "/admin/post/")
	if postID == "" {
		log.Println("ID Param 'key' is missing")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println("Error")
		http.Error(w, err.Error(), 400)
		return
	}
	db.DB.Model(&post).Where("ID = ?", postID).Updates(post)
	//json resp
	msg := Message{"success", "Post updated."}
	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Delete removes a post
func Delete(w http.ResponseWriter, r *http.Request) {
	var post Post
	postID := strings.TrimPrefix(r.URL.Path, "/admin/post/")
	if postID == "" {
		log.Println("ID Param 'key' is missing")
		return
	}
	// unmarshal content to ApplicantJSON
	db.DB.Find(&post, postID)
	db.DB.Delete(&post)
	//json resp
	msg := Message{"success", "The post has been deleted."}
	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
