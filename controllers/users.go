package controllers

import (
	"fmt"
	"net/http"

	"goglasses.com/m/v0/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// a new user account
//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the new method!!")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// This is used to process the signup from when a user tries to
// create a new user account
//POST Signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a fake message create a new user account")
}
