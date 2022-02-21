package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}


func main()  {
	var c Circle
	fmt.Println(c.radius)
	c.radius = 10.00
	fmt.Println(c.getArea())
	c.changeRadius(20)
	fmt.Println(c.radius)
	change(&c, 30)
	fmt.Println(c.radius)
}
func (c Circle) getArea() float64  {
	return c.radius * c.radius
}
// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64)  {
	c.radius = radius
}

// 以下操作将不生效
//func (c Circle) changeRadius(radius float64)  {
//   c.radius = radius
//}
// 引用类型要想改变值需要传指针
func change(c *Circle, radius float64)  {
	c.radius = radius
}
