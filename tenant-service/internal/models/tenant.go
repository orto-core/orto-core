package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name   string `json:"fullname"`
	Status string `json:"status"`
}

var Models []interface{} = []interface{}{Tenant{}}
