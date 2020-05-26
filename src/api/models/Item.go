package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	First_Name  string
	Last_Name   string
	Description string
}
