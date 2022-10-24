package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func busi(i int) {

	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	wg.Done()
}

func main() {
	//模拟用户需求业务的数量
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		go busi(i)
	}

	wg.Wait()
}
