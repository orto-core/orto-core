package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
	Status   string `gorm:"default:inactive"`
}

var Models []interface{} = []interface{}{User{}}
