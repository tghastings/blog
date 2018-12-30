package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/tghastings/blog/admin/db"
)

// Post describes the database schema
type Post struct {
	gorm.Model
	Title   string
	Date    string
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
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(post.ID)
	db.DB.Find(&post, post.ID)
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
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(post)
	db.DB.Model(&post).Where("ID = ?", post.ID).Updates(post)
}

// Delete removed a post
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
