package main

import "fmt"

//https://studygolang.com/articles/12745?fr=sidebar
func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum1(s[:len(s)/2], c)
	go sum1(s[len(s)/2:], c)
	x, y, z := <-c, <-c, <-c // receive from c
	fmt.Println(x, y, z, x+y+z)
}
