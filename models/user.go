package models

import "gorm.io/gorm"

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `jsong:"email" gorm:"unique"`
	Password      []byte `json:"-"`
	MessageThread []MessageThread
	gorm.Model
}
