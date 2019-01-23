package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	UserId int `gorm:"index"`
	Name   string
}
