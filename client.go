package database

import (
	"log"
	"github.com/jinzhu/gorm"

)
//connector variable
var Connector *gorm.DB


//creating my sq1 connection

func Connect(connectionString String) error {
	var err error
	Connector,err  = gorm.Open("mysql",connectionString)

	if err != nil {
		return err
}
log.Println("Connectipm Success")
return nil
}
