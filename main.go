package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct{
	Name string `json:"name"`
	Email string `json:"user_email"`
 	Password string `jason:"password"`
}
type MyInfo struct{
	Id string `json:"id"`
	Name string `json:"name"`

}

type Post struct{
	Userid string `json:"id"`
	Name string `json:"name"`
	Familyname string `json:"familyname"`
	Content string `json:"content"`
	Postid string `json:"postid"`
	Date string `json:"date"`
}

var db *sql.DB
var err error

func main(){
	fmt.Println("go-test")

	db, err = sql.Open("mysql", "root:@/facebookdb")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Connected")

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("POST")
	router.HandleFunc("/users", getUsersById).Methods("POST")
	router.HandleFunc("/createposts", createPost).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	//router.HandleFunc("/getposts/{id}", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT first_name,email,password FROM users")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var user User
	  err := result.Scan(&user.Name,&user.Email,&user.Password)
	  if err != nil {
		panic(err.Error())
	  }
	  users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)

  }

  func getUsersById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars((r))
	result, err := db.Query("SELECT first_name,email,password FROM users WHERE ID = ?", params["ID"] )
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {

	  err := result.Scan(&user.Name,&user.Email,&user.Password)
	  if err != nil {
		panic(err.Error())
	  }

	}
	json.NewEncoder(w).Encode(user)

  }

  func createPost(w http.ResponseWriter, r *http.Request) {  
stmt, err := db.Prepare("INSERT INTO post(Users_ID,content_of_Post) VALUES(?,?)")
  if err != nil {
    panic(err.Error())
  }  
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }  
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  ID := keyVal["user_id"]
  content := keyVal["content"]
  
  _, err = stmt.Exec(ID,content)

  if err != nil {
    panic(err.Error())
  }  
  fmt.Fprintf(w, "New post was created")
}

func signUp(w http.ResponseWriter, r *http.Request) {  
	stmt, err := db.Prepare("INSERT INTO users(first_name,family_name,email,password,Phone_number) VALUES(?,?,?,?,?)")
	  if err != nil {
		panic(err.Error())
	  }  
	  body, err := ioutil.ReadAll(r.Body)
	  if err != nil {
		panic(err.Error())
	  }  
	  keyVal := make(map[string]string)
	  json.Unmarshal(body, &keyVal)
	  firstname := keyVal["first_name"]
	  familyname := keyVal["family_name"]
	  email := keyVal["email"]
	  password := keyVal["password"]
	  phonenumber := keyVal["phone_number"]
	  
	  _, err = stmt.Exec(firstname, familyname, email, password,phonenumber)
	
	  if err != nil {
		panic(err.Error())
	  }  
	  fmt.Fprintf(w, "New user has signed in")
	} 


	func login(w http.ResponseWriter, r *http.Request) {  
		result, err := db.Prepare("SELECT ID, first_name FROM users WHERE email=? AND password=?")
		  if err != nil {
			panic(err.Error())
		  }  
		  body, err := ioutil.ReadAll(r.Body)
		  if err != nil {
			panic(err.Error())
		  }  
		  keyVal := make(map[string]string)
		  json.Unmarshal(body, &keyVal)
		  email := keyVal["email"]
		  password := keyVal["password"]


		  _, err = result.Exec(email, password)
		 // var info MyInfo
		//   var info MyInfo
		//   for result.Next() {
	  
		// 	err := result.Scan(&info.Id)
		// 	if err != nil {
		// 	  panic(err.Error())
		// 	}
	  
		//   }
		//   json.NewEncoder(w).Encode(id)

		} 
// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)  
// 	result, err := db.Query("SELECT users.ID, users.first_name, users.family_name, post.ID,post.content_of_Post, post.Date_of_Post FROM post Inner Join friends On post.Users_ID = friends.Users_recieved_ID Inner Join users On post.Users_ID = users.ID Where friends.Status_of_Request = 'confirmed' And friends.Users_sent_ID = ? UNION SELECT users.ID, users.first_name, users.family_name, post.ID,post.content_of_Post, post.Date_of_Post FROM post Inner Join friends On post.Users_ID = friends.Users_sent_ID Inner Join users On post.Users_ID = users.ID Where friends.Status_of_Request = 'confirmed' And friends.Users_recieved_ID = ? UNION SELECT users.ID, users.first_name, users.family_name, post.ID,post.content_of_Post, post.Date_of_Post FROM post Inner Join users On post.Users_ID = users.ID Where  users.ID = ? ORDER BY Date_of_Post", params["ID"],params["ID"],params["ID"])
// 	if err != nil {
// 	  panic(err.Error())
// 	}  
// 	defer result.Close() 

// 	var post Post  
// 	for result.Next() {
// 	  err := result.Scan(&post.Userid, &post.Name,&post.Familyname, &post.Postid, &post.Content, &post.Date)
// 	  if err != nil {
// 		panic(err.Error())
// 	  }
// 	}  
// 	json.NewEncoder(w).Encode(post)
//   }