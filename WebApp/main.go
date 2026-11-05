// Download and add gorm.io/gorm, gorm.io/driver/mysql, and github.com/gorilla/mux as requirements in go.mod
// Generate go.sum: 

// go get gorm.io/gorm gorm.io/driver/sqlite github.com/gorilla/mux

package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func initializeRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	InitialMigration()
	// log.Println("migration done")
	initializeRouter()
}