package model

import (
	"github.com/deliangyang/tt2/conections"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UserId int `gorm:"index"`
	Name   string
}

func GetUserInfoByUserId(userId int) (*Product, error) {
	product := &Product{}
	d := conections.DB.Where("user_id = ?", userId).First(product)
	return product, d.Error
}