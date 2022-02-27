package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
//query参数
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8000/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run()
}
