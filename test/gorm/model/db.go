package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"sync"
	"time"
)

var DB *gorm.DB
var once sync.Once

type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level           = values[0]
		source          = values[1]
	)
	if level == "sql" {
		sql := values[3].(string)
		log.Println(sql, level, source)
	} else {
		log.Println(values)
	}
}

func GetInstance() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := "root:123456@/test?charset=utf8&parseTime=True&loc=Local"
		if len(dsn) == 0 {
			log.Fatal("Failed to get database dsn")
		}
		DB, err = gorm.Open("mysql", dsn)
		DB.LogMode(true)
		DB.SetLogger(&MyLogger{})
		if err != nil {
			log.Fatalf("open db: %s", err)
		}

		sqlDb := DB.DB()
		sqlDb.SetMaxOpenConns(100)
		sqlDb.SetMaxIdleConns(100)
		sqlDb.SetConnMaxLifetime(time.Hour * 8)
		//DB.SingularTable(true)//自动为对象名加s user=>users
	})
	return DB
}

