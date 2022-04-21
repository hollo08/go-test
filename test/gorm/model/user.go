package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)



type User struct {
	gorm.Model
	Name string
	Country string
	Age sql.NullInt16
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
