package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type aaa uint8

const (
	a      = 1
	b      = 2
	aa aaa = iota
	bb
	cc
	dd
	ee
	ff
)

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(aa)
	fmt.Println(bb)
	var a int64
	print(a)
}
