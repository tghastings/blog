package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "log"
  "github.com/tghastings/blog/admin/db"
  "github.com/tghastings/blog/admin/users"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte("captainjacksparrowsayshi")

var err error

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type Message struct {
  Type    string
  Message string
}

type User struct {
  gorm.Model
  Username string
  Password string
  APIToken string
  Valtoken string
  Email string
}

func main() {
  if err := db.Open(); err != nil {
    // handle error
  }
  defer db.Close()
  // Migrate the schema
  db.DB.AutoMigrate(&Product{})
  db.DB.AutoMigrate(user.User{})
  
  var userCount int
  var users []User
  db.DB.Find(&users).Count(&userCount);
  if (userCount == 0) {
    db.DB.Create(&User{Username: "root", Password: "$2a$04$7ZZOLkODB70E5UL9UqvGzuPnqfaCZjKVUd7UhYP4jRywU/gOzHomS"})
    fmt.Println("Created new user `root` password is `12345`")
  }


  http.HandleFunc("/", Index)
  http.Handle("/admin/users", isAuthorized(user.Index))
  http.HandleFunc("/show", Show)
  http.HandleFunc("/auth", user.Auth)
  http.HandleFunc("/admin/user", user.Show)
  http.HandleFunc("/create", Create)
  http.HandleFunc("/admin/user/create", user.Create)
  http.HandleFunc("/delete", Delete)
  http.HandleFunc("/admin/user/delete", user.Delete)
  http.HandleFunc("/update", Update)
  http.HandleFunc("/admin/user/update", user.Update)
  
  fmt.Println("The application has started and is listening on port 8090.")

  http.ListenAndServe(":8090", nil)
  log.Fatal(http.ListenAndServe(":8090", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {

  var products []Product
  db.DB.Find(&products)
  js, err := json.Marshal(products)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Create(w http.ResponseWriter, r *http.Request) {
  var newProduct Product
  err := json.NewDecoder(r.Body).Decode(&newProduct)
  if err != nil {
    // handle error
    log.Fatal("Error")
  }
  if err != nil {
    log.Fatal(err)  
  }
  // Create
  db.DB.Create(&newProduct)

  //json resp
  msg := Message{"success", "Product added."}
  js, err := json.Marshal(msg)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Show(w http.ResponseWriter, r *http.Request) {
  var product Product
  err := json.NewDecoder(r.Body).Decode(&product)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  fmt.Println(product.ID)
  db.DB.Find(&product, product.ID)
  js, err := json.Marshal(&product)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Update(w http.ResponseWriter, r *http.Request) { 
  var product Product
  err := json.NewDecoder(r.Body).Decode(&product)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  fmt.Println(product)
  db.DB.Model(&product).Where("ID = ?", product.ID).Updates(product)
}

func Delete(w http.ResponseWriter, r *http.Request) { 
  var product Product
  err := json.NewDecoder(r.Body).Decode(&product)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  fmt.Println(product.ID)
  // unmarshal content to ApplicantJSON
  db.DB.Find(&product, product.ID)
  db.DB.Delete(&product)
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