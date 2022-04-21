package main

import (
	"github.com/hollo08/go-test/test/gorm/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main()  {
	db := model.GetInstance()
	defer db.Close()
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
}