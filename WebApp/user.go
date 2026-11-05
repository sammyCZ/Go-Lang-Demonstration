// GET all users
// curl http://localhost:8080/users

// GET a single user 
// curl http://localhost:8080/users/#

// PUT (update) a user
// curl -X PUT http://localhost:8080/users/# \
//   -H "Content-Type: application/json" \
//   -d '{"firstname":"Jane","lastname":"Doe","email":"janet@example.com"}'

// DELETE a user
// curl -X DELETE http://localhost:8080/users/#

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"runtime"
)

var DB *gorm.DB

type User struct {
	// Add a pre-defined struct (CreatedAt, UpdatedAt, DeletedAt)
	gorm.Model
	// Send data with json 
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
}

// Initialize and connect to sqlite database 
func InitialMigration() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&User{})
}

// -------------------------- Define Handler Methods

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Garbage collection! Go allocates memory for this.
	// Once the handler returns and the HTTP is written,
	// nothing references that memory anymore
	var users []User
	runtime.GC() 
	
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is Deleted")
}