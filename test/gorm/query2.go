package main

import (
	"fmt"
	"github.com/hollo08/go-test/test/gorm/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	find()
	take()
	last()
	find()
	OtherFirst()
	where()
	where2()
}

func first(){
	db := model.GetInstance()
	user := &model.User{}
	db.First(user)
	fmt.Printf("user value is %v \n", user)
}

func take(){
	db := model.GetInstance()
	user := &model.User{}
	db.Take(user)
	fmt.Printf("user value is %v \n", user)
}

func last(){
	db := model.GetInstance()
	user := &model.User{}
	db.Last(user)
	fmt.Printf("user value is %v \n", user)
}

func find(){
	db := model.GetInstance()
	user := &model.User{}
	db.Find(user)
	fmt.Printf("user value is %v \n", user)
}

func OtherFirst(){
	db := model.GetInstance()
	user := &model.User{}
	db.First(user, 10)
	fmt.Printf("user value is %v \n", user)
}

func where(){
	db := model.GetInstance()
	users := &[]model.User{}
	db.Where("name in (?) and age>= ?", []string{"aaa", "bbb"}, 0).Find(users)
	fmt.Printf("user value is %v \n", users)
}

func where2(){
	db := model.GetInstance()
	users := make([]model.User, 0)
	db.Where(map[string]interface{}{"name": "aaa", "age": 2}).Find(&users)
}

