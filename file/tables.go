package file

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

var Db *gorm.DB
var err error
var once sync.Once
func GetDb() {
	once.Do(func() {
		Db, err = gorm.Open("sqlite3", "C:\\files.db")
		if err != nil {
			panic(err)
		}
	})
}

type Files struct {
	gorm.Model
	Name string `gorm:"index"`
	Path string `gorm:"index"`
	FileType string`gorm:"index"`
	DirModTime time.Time
}
