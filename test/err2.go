package main

import (
	"fmt"
)

// 自定义错误信息结构
type DIV_ERR struct {
	etype int  // 错误类型
	v1 int     // 记录下出错时的除数、被除数
	v2 int
}
// 实现接口方法 error.Error()
func (div_err DIV_ERR) Error() string {
	fmt.Println(div_err.v1, div_err.v2)
	if 0==div_err.etype {
		return "除零错误"
	}else{
		return "其他未知错误"
	}
}
// 除法
func div(a int, b int) (int,*DIV_ERR) {
	if b == 0 {
		// 返回错误信息
		return 0,&DIV_ERR{0,a,b}
	} else {
		// 返回正确的商
		return a / b, nil
	}
}
func main() {
	// 正确调用
	v,r :=div(100,2)
	if nil!=r{
		fmt.Println("(1)fail:",r)
	}else{
		fmt.Println("(1)succeed:",v)
	}
	// 错误调用
	v,r =div(100,0)
	if nil!=r{
		fmt.Println("(2)fail:",r)
	}else{
		fmt.Println("(2)succeed:",v)
	}
}
