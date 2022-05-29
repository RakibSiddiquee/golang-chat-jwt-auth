package database

import (
	"github.com/RakibSiddiquee/golang-chat-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/chat_app"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}
