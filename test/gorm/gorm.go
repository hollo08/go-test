package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&Product{})

	// 增
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 查
	var product Product
	db.First(&product, 1) // 找到id为1的产品
	db.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品

	// 改 - 更新产品的价格为 2000
	db.Model(&product).Update("Price", 2000)

	// 删 - 删除产品
	db.Delete(&product)
}
