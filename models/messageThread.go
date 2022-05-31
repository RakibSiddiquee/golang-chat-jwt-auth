package models

import "gorm.io/gorm"

type MessageThread struct {
	Id      uint `json:"id"`
	UserId  uint
	Message []Message
	gorm.Model
}
