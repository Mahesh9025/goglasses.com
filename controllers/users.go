package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"goglasses.com/m/v0/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
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
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
