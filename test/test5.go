package main

import "fmt"

var c int = 5

func main(){
	var a, b int
	a = 20
	b = 30
	//change c
	c = 3
	fmt.Printf("a=%d, b=%d, c=%d, a+b=%d \n", a, b, c, a+b)
	testScope()
	fmt.Printf("c=%d \n", c)
	testArray()
}
//change c
func testScope(){
	c = 4
}
func testArray(){
	balance := [5]int{1:3, 3:4}
	//balance[5] = 4
	fmt.Printf("value of index 0 is %d \n", balance[0])
	fmt.Printf("value of index 1 is %d \n", balance[1])
	for i:=0; i<len(balance); i++ {
		fmt.Printf("value of index %d is %d \n", i, balance[i])
	}
	mySlice := balance[1:]
	fmt.Println(mySlice)
}
