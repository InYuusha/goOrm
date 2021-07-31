package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
	defer db.Close() //if panicked ,
	// do this before exit(LIFO order for multiple defer)
	db.AutoMigrate(&User{}) //migrate schema ,create tables update table ,(wont del)
}
func AllUsers(res http.ResponseWriter, req *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Couldnt connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(res).Encode(users)
}

func NewUser(res http.ResponseWriter, req *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Couldnt connect to database")
	}
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(res, "New User Created")
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Couldnt connect to database")
	}
	defer db.Close()
	vars := mux.Vars(req)
	name := vars["name"]

	var user User
	db.Where("name =?", name).Find(&user) //populates the result in user obj
	db.Delete(&user)

	fmt.Fprintf(res, "User Successfully Deleted")
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Couldnt connect to database")
	}
	defer db.Close()
	vars:=mux.Vars(req)

	name:=vars["name"]
	email:=vars["email"]
	
	var user User
	db.Where("name =?",name).Find(&user)
	user.Email=email;
	db.Save(&user)

	fmt.Fprintf(res,"Successfully Updated User")
}
