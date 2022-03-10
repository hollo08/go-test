package main

import "fmt"

type LinkType uint8
const (
    ShareClick LinkType = 1 //共享流量
    OriginClick		LinkType = 2		//原文跳转
)

func main(){
    fmt.Println(OriginClick)
}


