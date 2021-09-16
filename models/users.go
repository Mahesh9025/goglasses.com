package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//ErrNotFound is returned when a resource cannot be found
	// In the database
	ErrNotFound  = errors.New("models: resource not found")
	ErrInvalidID = errors.New("models: ID must be > 0.")
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	fmt.Println("database connected")
	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// First will query using provided by gorm.DB and it will
// get first item returned and place into dst. If
// nothing is found in the query , it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// create a user provided by user data
// like Id, created_at, updated_at, name and email
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
	// return nil
}

// update a user provided by user data
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
	// return nil
}

// Delete a user provided by userID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// closes the userService database connection
func (us *UserService) close() error {
	return us.db.Close()
}

// DestructiveReset drops the user table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
