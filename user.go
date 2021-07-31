package main

import(
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct{
	gorm.Model
	Name string
	Email string
}
func InitialMigration(){
	db,err = gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
	defer db.close() //if panicked ,
	// do this before exit(LIFO order for multiple defer)

	db.AutoMigrate(&User{}) //migrate schema ,create tables update table ,(wont del)
}

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