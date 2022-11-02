package main

import (
	"fmt"
	"sort"
)

func main() {
	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)
	//GuessingGame()
	SearchNum()
	calCircle()
	fmt.Println("a3")
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func SearchNum() {
	x := 10
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
}

func calCircle() {
	fmt.Printf("a - b = %f \n", 0.5/3.14)
}
