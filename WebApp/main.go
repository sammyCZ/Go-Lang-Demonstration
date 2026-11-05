// Download and add gorm.io/gorm, gorm.io/driver/mysql, and github.com/gorilla/mux as requirements in go.mod
// Generate go.sum: 
// go get gorm.io/gorm gorm.io/driver/mysql github.com/gorilla/mux. 

package main

import "github.com/gorilla/mux"

func initializeRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}