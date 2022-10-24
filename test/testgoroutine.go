package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	//模拟用户需求业务的数量
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		go func(i int) {
			//... do some busi...

			fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}
