package main

import(
	"fmt"
	"net/http"
)

func AllUsers(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"All users Endpoint ")
}

func NewUser(res http.ResponseWriter,req*http.Request){
	fmt.Fprintf(res,"New User Endopint")
}

func DeleteUser(res http.ResponseWriter,req*http.Request){
	fmt.Fprintf(res,"Delete user Endpoint")
}

func UpdateUser(res http.ResponseWriter,req*http.Request){
	fmt.Fprintf(res,"Update user Endpoint")
}