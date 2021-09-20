package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"goglasses.com/m/v0/models"
	"goglasses.com/m/v0/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
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
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// POST/Login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email address")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password provided")
	case nil:
		fmt.Println(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
