package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//go say("world")
	//say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 2)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	time.Sleep(100 * time.Millisecond)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
	e := make(chan int, 10)
	go fibonacci2(cap(e), e)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range e {
		fmt.Println(i)
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}