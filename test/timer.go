package main

import (
    "reflect"
    "time"
)
import "fmt"

func main(){
    parse1()
    parse2()
    parse3()
}

func parse1(){
    timestamp := 1602482956
    fmt.Println(reflect.TypeOf(timestamp))
    timeobj := time.Unix(int64(timestamp), 0)
    date := timeobj.Format("2006-01-02 15:04:05")//06-01-02 15:04:05, 2006/01/02 15:04:05, 2006/01/02, 2006-01-02
    fmt.Println(date)
}
func parse2(){
    str := "2020-10-12 14:09:17"
    tmp := "2006-01-02 15:04:05"
    res, _ := time.ParseInLocation(tmp, str, time.Local)
    fmt.Println(res.Unix())
}
func parse3(){
    fmt.Println(time.Millisecond)
    fmt.Println(time.Now())
    t := time.Now().Unix()
    fmt.Println(reflect.TypeOf(t))
    fmt.Println(time.Unix(t, 0).Format("2006-01-02 15:04:05"))
    t += 86400 //增加一天
    fmt.Println(time.Unix(t, 0).Format("2006-01-02 15:04:05"))
}

