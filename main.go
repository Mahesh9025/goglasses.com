package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"goglasses.com/m/v0/controllers"
	"goglasses.com/m/v0/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "goglasses_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	must(err)
	// defer us.Close()
	// us.DestructiveReset()
	us.AutoMigrate()
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/contact", staticController.Contact).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	fmt.Println("Starting the server on 8525")
	http.ListenAndServe(":8525", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
