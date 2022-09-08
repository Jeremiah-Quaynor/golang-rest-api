package main

import (
	"golang-rest-api"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main () {
	config := 
		database.Config{
			ServerName: "localhost:3306",
			User: "root",
			Password: "root",
			DB: "learning",
		}

		connectionString := database.GetConnectionString(config)
		err := database.Connect(connectionString)
		if err != nil {
			panic(err.Error())
		}
}