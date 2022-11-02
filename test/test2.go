package main

import (
	"fmt"
	"strings"
)

func main() {
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				fmt.Printf("%d  不是素数\n", i)
				break // 如果发现因子，则不是素数
			}

			break
		}
		//if(j > (i/j)) {
		fmt.Printf("%d  是素数\n", i)
		//}
	}
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println("a2")
}
