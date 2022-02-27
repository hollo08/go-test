package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//url参数
func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action2 := strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action + " is "+action2)
	})
	//默认为监听8080端口
	r.Run(":8000")
}
