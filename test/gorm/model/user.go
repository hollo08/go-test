package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Country string
	Age int32
}

type Product struct {
	gorm.Model
	Name string
	Barcode string
	Price float64
	Desc string
}

type Order struct {
	gorm.Model
	UserId int64
	ProductId int64
	Price float64
}
