package main

import (
	"fmt"
	model "github.com/hollo08/go-test/test/gorm/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	query1()
}



func query1(){
	db := model.GetInstance()
	defer db.Close()
	user := &model.User{}
	db.Debug().First(user)
	fmt.Printf("user value is %v", user)
}
