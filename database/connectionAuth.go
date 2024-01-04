package database

import (
	"github.com/CVWO/go-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBAuth *gorm.DB

func ConnectToDb() {
	connection, err := gorm.Open(mysql.Open("root:Password1107@/alternate_go_crud"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DBAuth = connection

	connection.AutoMigrate(&models.User{})
}
