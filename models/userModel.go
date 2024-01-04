package models

type User struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Username string `gorm:"unique"`
	Password string
}
