package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"goglasses.com/m/v0/controllers"
)

func main() {
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/contact", staticController.Contact).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	http.ListenAndServe(":8525", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
