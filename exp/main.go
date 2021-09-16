package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	if err != nil {
		panic(err)
	}

	// defer us.Close()
	us.DestructiveReset()
	user := models.User{
		Name:  "Mahesh",
		Email: "mahesh90@yopmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	user.Email = "mahesh@gmail.com"
	if err := us.Update(&user); err != nil {
		panic(err)
	}
	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)

	userByEmail, err := us.ByEmail(user.Email)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByEmail)

	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	fmt.Println("Deleted")
}
