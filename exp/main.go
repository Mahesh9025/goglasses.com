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
	// us.DestructiveReset()
	user, err := us.ByID(1)
	fmt.Println(user, err)

}
