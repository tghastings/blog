package user

import (
  "fmt"
  "encoding/json"
  "net/http"
  "log"
  "github.com/jinzhu/gorm"
  "crypto/rand"
  "golang.org/x/crypto/bcrypt"
  "github.com/tghastings/blog/admin/db"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  jwt "github.com/dgrijalva/jwt-go"
  "time"
)

var mySigningKey = []byte("pleasedonthackmebro")

type Message struct {
  Type    string
  Message string
}

type APIToken struct {
  APIToken string
}

type JSONTOKEN struct {
  Token    string
}

type User struct {
  gorm.Model
  Username string
  Password string
  APIToken string
  Email string
}

type Cookie struct {
  Name       string
  Value      string
  Path       string
  Domain     string
  Expires    time.Time
  RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
  MaxAge   int
  Secure   bool
  HttpOnly bool
  Raw      string
  Unparsed []string // Raw text of unparsed attribute-value pairs
}

func Index(w http.ResponseWriter, r *http.Request) {
  var users []User
  db.DB.Find(&users)
  js, err := json.Marshal(users)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Create(w http.ResponseWriter, r *http.Request) {
  var newUser User
  err := json.NewDecoder(r.Body).Decode(&newUser)
  if err != nil {
    log.Fatal(err)  
  }
  // Make token
  token := tokenGenerator()
  JSONtoken := APIToken{token}
  js, err := json.Marshal(JSONtoken)
  if err != nil {
    // handle error
    log.Fatal(err)
  }
  fmt.Println(newUser.Password)
  newUser.Password = hashAndSalt([]byte(newUser.Password))
  fmt.Println(newUser.Password)
  err = json.Unmarshal(js, &newUser)
  if err != nil {
    // handle error
    log.Fatal(err)
  }
  // Create
  db.DB.Create(&newUser)

  //json resp
  msg := Message{"success", "User added."}
  js, err = json.Marshal(msg)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Show(w http.ResponseWriter, r *http.Request) {
  var user User
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  fmt.Println(user.ID)
  db.DB.Find(&user, user.ID)
  js, err := json.Marshal(&user)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func Update(w http.ResponseWriter, r *http.Request) {
  var user User
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  // fmt.Println(user)
  db.DB.Model(&user).Where("ID = ?", user.ID).Updates(user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
  var user User
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  // fmt.Println(user.ID)
  // unmarshal content to ApplicantJSON
  db.DB.Find(&user, user.ID)
  db.DB.Delete(&user)
}

func Auth(w http.ResponseWriter, r *http.Request) {
  var user User
  fmt.Printf("%+v", r)
  // var count int
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  // fmt.Println(user.Username)
  plainPassword := user.Password
  // fmt.Println(plainPassword)
  // unmarshal content to ApplicantJSON
  db.DB.Where("username = ?", user.Username).Find(&user)
  // fmt.Println(user.Password)

  // do the passwords match?
  if comparePasswords(user.Password, []byte(plainPassword)) {
    //json resp
    newToken := GenerateJWT(user.Username)
    // JWT := JSONTOKEN{newToken}
    // js, err := json.Marshal(JWT)
    // if err != nil {
    //   log.Println(err)
    // }
    // w.Header().Set("Content-Type", "application/json")
    // w.Write(js)
    
    expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{Name: "username", Value: user.Username, Expires: expiration}
    http.SetCookie(w, &cookie)
    cookie = http.Cookie{Name: "Token", Value: newToken, Expires: expiration}
    http.SetCookie(w, &cookie)

  } else {
    fmt.Fprintf(w, "Bad username and/or password")
  }
}

func hashAndSalt(pwd []byte) string {
    
  // Use GenerateFromPassword to hash & salt pwd.
  // MinCost is just an integer constant provided by the bcrypt
  // package along with DefaultCost & MaxCost. 
  // The cost can be any value you want provided it isn't lower
  // than the MinCost (4)
  hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
  if err != nil {
      log.Println(err)
  }
  // GenerateFromPassword returns a byte slice so we need to
  // convert the bytes to a string and return it
  return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
  // Since we'll be getting the hashed password from the DB it
  // will be a string so we'll need to convert it to a byte slice
  byteHash := []byte(hashedPwd)
  err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
  if err != nil {
      log.Println(err)
      return false
  }

  return true
}

func GenerateJWT(username string) string {
  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)

  claims["authorized"] = true
  claims["client"] = username

  tokenString, err := token.SignedString(mySigningKey)

  if err != nil {
    fmt.Errorf("Something Went Wrong: %s", err.Error())
    return "Error, unable to make JWT in user.go"
  }
  //json resp
  return tokenString
}

func tokenGenerator() string {
  b := make([]byte, 32)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}