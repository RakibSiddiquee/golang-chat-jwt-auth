package models

import "gorm.io/gorm"

type Message struct {
	Id              uint `json:"id"`
	MessageThreadId uint `json:"message_thread_id"`
	FromUserId      uint
	ToUserId        uint
	Message         string `json:"message"`
	FromUser        User   `gorm:"foreignKey:FromUserId"`
	ToUser          User   `gorm:"foreignKey:ToUserId"`
	gorm.Model
}
