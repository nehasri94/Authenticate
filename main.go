package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mux"
	"net/http"
	"database/sql"
	"mysql"
)

var uname,email,password string

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "Neta"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func loginpagehandler(w http.ResponseWriter, r *http.Request){

	content, err := ioutil.ReadFile("login.html")
	if err != nil {
		log.Fatal(err)
	}
		fmt.Fprint(w,string(content))
}
func loginhandler(w http.ResponseWriter, r *http.Request){

	uname := r.FormValue("uname")
	password := r.FormValue("password")
	if len(uname)!=0 && len(password)!=0 && uname=="neha" && password=="neha" {
		//http.Redirect(w,r,"/",http.StatusMovedPermanently)
		fmt.Fprint(w,"logged in")
	}else if uname!="neha" && len(uname)!=0 && len(password)!=0 {
		http.Redirect(w,r,"/register",http.StatusMovedPermanently)
	}else {
		http.Redirect(w,r,"/",http.StatusMovedPermanently)
	}

}
func registerpagehandler(w http.ResponseWriter, r *http.Request){
	content, err := ioutil.ReadFile("register.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w,string(content))

}

func registerhandler(w http.ResponseWriter, r *http.Request){
	uname := r.FormValue("uname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if len(uname)!=0 && len(email)!=0 && len(password)!=0 {
		http.Redirect(w,r,"/",http.StatusMovedPermanently)
	}else {
		http.Redirect(w,r,"/register",http.StatusMovedPermanently)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/",loginpagehandler).Methods("GET")
	router.HandleFunc("/login",loginhandler).Methods("POST")
	router.HandleFunc("/register",registerpagehandler).Methods("GET")
	router.HandleFunc("/register",registerhandler).Methods("POST")
	http.ListenAndServe(":8000", router)

}