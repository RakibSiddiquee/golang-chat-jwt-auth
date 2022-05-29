package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `jsong:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
