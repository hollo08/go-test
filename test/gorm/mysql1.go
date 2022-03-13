package main

import (
	"fmt"
	model "github.com/hollo08/go-test/test/gorm/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	mirgrate()
	//query1()
}

func mirgrate(){
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {

	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
}

func query1(){
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	user := &model.User{}
	if err != nil {

	}
	db.Debug().First(user)
	fmt.Printf("user value is %v", user)
}
