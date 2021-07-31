package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "HEllo World")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", helloWorld).Methods("GET")

	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	fmt.Println("Starting...")

	InitialMigration()//setup table

	handleRequest()
}
