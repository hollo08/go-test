package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//Person ..
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

//http://localhost:8080/5lmh?age=11&name=枯藤&birthday=2016-01-02
func main() {
	r := gin.Default()
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
