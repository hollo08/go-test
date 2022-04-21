package main

import (
	"fmt"
	"github.com/hollo08/go-test/test/gorm/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	groupquery()
	//zeroquery()
}

func groupquery(){
	type group struct {
		date int64
		total int32
	}
	groups := make([]*model.User, 0)
	db := model.GetInstance()
	err := db.Debug().Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(age)").Having("sum(age) > ?", 10).Find(groups).Error
	if err != nil {

	}
	for _, g := range groups{
		fmt.Printf("group value is %v \n", g)
	}
}

func zeroquery(){

}